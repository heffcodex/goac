package goacoap

type IPermGetter interface {
	Permissions() []Permission
}

type INode interface {
	IPermGetter
	Name() string
}

type IParametrized interface {
	INode
	Allowed() bool
	Param(name string) IParameter
}
