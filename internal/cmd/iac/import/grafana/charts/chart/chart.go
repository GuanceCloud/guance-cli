package chart

type BuildOptions struct {
	Group       string
	Measurement string
}

type Builder interface {
	Build(m map[string]any, opts BuildOptions) (map[string]any, error)
}
