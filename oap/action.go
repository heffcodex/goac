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
	Parameter
}

func (a *Action) SetAllow(v bool) IAction { a.Parameter.SetAllow(v); return a }
func (a *Action) Allow() IAction          { a.Parameter.Allow(); return a }
func (a *Action) Deny() IAction           { a.Parameter.Deny(); return a }

func (a *Action) End() IAction {
	a.Parameter.End()
	return a
}

func (a *Action) Finalize() IAction {
	a.Parameter.Finalize()
	return a
}
