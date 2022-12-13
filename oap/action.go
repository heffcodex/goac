package goacoap

type IAction interface {
	IParametrized
	SetAllow(v bool) *Action
	Allow() *Action
	Deny() *Action
	End() *Action
	Finalize() *Action
}

var _ IAction = (*Action)(nil)

type Action struct {
	Param
}

func (a *Action) SetAllow(v bool) *Action { a.Param.SetAllow(v); return a }
func (a *Action) Allow() *Action          { a.Param.Allow(); return a }
func (a *Action) Deny() *Action           { a.Param.Deny(); return a }

func (a *Action) End() *Action {
	a.Param.End()
	return a
}

func (a *Action) Finalize() *Action {
	a.Param.Finalize()
	return a
}
