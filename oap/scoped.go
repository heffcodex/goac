package goacoap

import "context"

type CtxProvider func() context.Context

type IScoped[O IObject, S any] interface {
	Base() O
	Scope() S
	Ctx() context.Context
	Err() error
	SetErr(err error) bool
}

var _ IScoped[IObject, any] = (*Scoped[IObject, any])(nil)

type Scoped[O IObject, S any] struct {
	base        O
	scope       S
	ctxProvider CtxProvider
	err         error
}

func NewScoped[O IObject, S any](base O, scope S, ctxProvider CtxProvider) *Scoped[O, S] {
	if ctxProvider == nil {
		ctxProvider = func() context.Context { return context.Background() }
	}

	return &Scoped[O, S]{
		base:        base,
		scope:       scope,
		ctxProvider: ctxProvider,
	}
}

func (s *Scoped[O, S]) Base() O              { return s.base }
func (s *Scoped[O, S]) Scope() S             { return s.scope }
func (s *Scoped[O, S]) Ctx() context.Context { return s.ctxProvider() }

func (s *Scoped[O, S]) Err() error            { return s.err }
func (s *Scoped[O, S]) SetErr(err error) bool { s.err = err; return s.err != nil }
