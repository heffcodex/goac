package oap

type IPathGetter interface {
	GetAllowedPaths() []Path
}

type INode interface {
	IPathGetter
	Name() string
}

type IParametrized interface {
	INode
	Allowed() bool
	Param(name string) *Param
}
