package promql

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

// NewRewriter will create new Rewriter
func NewRewriter(measurement string) *Rewriter {
	return &Rewriter{Measurement: measurement}
}

const magicInterval = "[999d]"

// Rewrite will rewrite a query from Grafana-specific PromQL to MetricsQL on Guance Cloud
func (w *Rewriter) Rewrite(query string) (string, error) {
	// Fix for Grafana-specific interval interpolation syntax
	// See also: https://docs.guance.com/dql/metricsql/
	query = strings.ReplaceAll(query, "[$__interval]", magicInterval)
	query = strings.ReplaceAll(query, "[$__rate_interval]", magicInterval)
	// Compatible for old-style grafana dashboard
	query = strings.ReplaceAll(query, "[$interval]", magicInterval)
	query = strings.ReplaceAll(query, "[$rate_interval]", magicInterval)

	expr, err := parser.ParseExpr(query)
	if err != nil {
		return "", fmt.Errorf("failed to parse query %q: %w", query, err)
	}
	if err := parser.Walk(w, expr, nil); err != nil {
		return "", fmt.Errorf("walk PromQL expression failed: %w", err)
	}
	return strings.ReplaceAll(parser.Prettify(expr), magicInterval, ""), nil
}

func (w *Rewriter) Visit(node parser.Node, path []parser.Node) (v parser.Visitor, err error) {
	if e, ok := node.(*parser.VectorSelector); ok {
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
	}
	return w, nil
}

var varPattern = regexp.MustCompile(`\$\w+`)

func (w *Rewriter) rewriteVar(name string) string {
	return varPattern.ReplaceAllStringFunc(name, func(s string) string {
		return fmt.Sprintf("#{%s}", s[1:])
	})
}

const nameSep = "_"

func (w *Rewriter) rewriteName(name string) string {
	measurement := w.Measurement
	if measurement == "" {
		tokens := strings.Split(name, nameSep)
		measurement = tokens[0]
		name = strings.Join(tokens[1:], nameSep)
		return fmt.Sprintf("%s:%s", measurement, name)
	}
	// Escape `-` in measurement name
	measurement = strings.ReplaceAll(measurement, "-", "\\-")
	return fmt.Sprintf("%s:%s", measurement, name)
}
