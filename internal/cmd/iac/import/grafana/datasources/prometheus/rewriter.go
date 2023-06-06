package prometheus

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/promql/parser"
)

// Rewriter is a rewriter to convert Grafana-specific PromQL to MetricsQL on Guance Cloud
type Rewriter struct {
	Measurement string
}

const magicInterval = "[999d]"

func (w *Rewriter) Rewrite(query string) (string, error) {
	// Fix for Grafana-specific interval interpolation syntax
	// See also: https://docs.guance.com/dql/metricsql/
	query = strings.ReplaceAll(query, "[$__interval]", magicInterval)
	query = strings.ReplaceAll(query, "[$__rate_interval]", magicInterval)

	expr, err := parser.ParseExpr(query)
	if err != nil {
		return "", fmt.Errorf("failed to parse query %q: %w", query, err)
	}

	return strings.ReplaceAll(parser.Prettify(w.rewriteVars(expr)), magicInterval, ""), nil
}

func (w *Rewriter) rewriteVars(expr parser.Expr) parser.Expr {
	switch e := expr.(type) {
	case *parser.VectorSelector:
		e.Name = w.rewriteName(e.Name)
		labelMatchers := make([]*labels.Matcher, 0)
		for i := 0; i < len(e.LabelMatchers); i++ {
			if e.LabelMatchers[i].Name == "__name__" {
				continue
			}
			e.LabelMatchers[i].Value = w.rewriteVar(e.LabelMatchers[i].Value)
			labelMatchers = append(labelMatchers, e.LabelMatchers[i])
		}
		e.LabelMatchers = labelMatchers
	case *parser.MatrixSelector:
		e.VectorSelector = w.rewriteVars(e.VectorSelector)
	case *parser.SubqueryExpr:
		e.Expr = w.rewriteVars(e.Expr)
	case *parser.BinaryExpr:
		e.LHS = w.rewriteVars(e.LHS)
		e.RHS = w.rewriteVars(e.RHS)
	case *parser.AggregateExpr:
		e.Expr = w.rewriteVars(e.Expr)
	case *parser.Call:
		for i, arg := range e.Args {
			e.Args[i] = w.rewriteVars(arg)
		}
	case *parser.ParenExpr:
		e.Expr = w.rewriteVars(e.Expr)
	case *parser.UnaryExpr:
		e.Expr = w.rewriteVars(e.Expr)
	}
	return expr
}

var varPattern = regexp.MustCompile(`\$\w+`)

func (w *Rewriter) rewriteVar(name string) string {
	return varPattern.ReplaceAllStringFunc(name, func(s string) string {
		return fmt.Sprintf("#{%s}", s[1:])
	})
}

func (w *Rewriter) rewriteName(name string) string {
	return fmt.Sprintf("%s:%s", w.Measurement, strings.TrimPrefix(name, fmt.Sprintf("%s_", w.Measurement)))
}
