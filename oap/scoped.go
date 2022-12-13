package goacoap

import "context"

type CtxProvider func() context.Context

type IScoped[T any] interface {
	Scope() T
	Ctx() context.Context
	Err() error
	SetErr(err error)
}

var _ IScoped[any] = (*Scoped[any])(nil)

type Scoped[T any] struct {
	scope       T
	ctxProvider CtxProvider
	err         error
}

func NewScoped[T any](scope T, ctxProvider CtxProvider) *Scoped[T] {
	if ctxProvider == nil {
		ctxProvider = func() context.Context { return context.Background() }
	}

	return &Scoped[T]{
		scope:       scope,
		ctxProvider: ctxProvider,
	}
}

func (s *Scoped[T]) Scope() T             { return s.scope }
func (s *Scoped[T]) Ctx() context.Context { return s.ctxProvider() }

func (s *Scoped[T]) Err() error       { return s.err }
func (s *Scoped[T]) SetErr(err error) { s.err = err }
