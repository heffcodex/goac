package goacoap

type IAction interface {
	IParametrized
	SetAllow(v bool) IAction
	Allow() IAction
	Deny() IAction
	End() IAction
	Finalize() IAction
}

var _ IAction = (*Action)(nil)

type Action struct {
	Param
}

func (a *Action) SetAllow(v bool) IAction { a.Param.SetAllow(v); return a }
func (a *Action) Allow() IAction          { a.Param.Allow(); return a }
func (a *Action) Deny() IAction           { a.Param.Deny(); return a }

func (a *Action) End() IAction {
	a.Param.End()
	return a
}

func (a *Action) Finalize() IAction {
	a.Param.Finalize()
	return a
}
