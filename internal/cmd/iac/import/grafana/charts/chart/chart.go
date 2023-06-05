package chart

type BuildOptions struct {
	Group string
}

type Builder interface {
	Build(m map[string]interface{}, opts BuildOptions) (map[string]interface{}, error)
}
