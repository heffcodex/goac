package oap

type IAction interface {
	IParametrized
	End() *Action
	Finalize() *Action
}

var _ IAction = (*Action)(nil)

type Action struct {
	Param
}

func (a *Action) End() *Action {
	a.Param.End()
	return a
}

func (a *Action) Finalize() *Action {
	a.Param.Finalize()
	return a
}
