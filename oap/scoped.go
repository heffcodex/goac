package goacoap

import "context"

type CtxProvider func() context.Context

type Scoped struct {
	base        IObject
	ctxProvider CtxProvider
	err         error
}

func NewScopedObject(base IObject, ctxProvider CtxProvider) *Scoped {
	if ctxProvider == nil {
		ctxProvider = func() context.Context { return context.Background() }
	}

	return &Scoped{
		base:        base,
		ctxProvider: ctxProvider,
	}
}

func (s *Scoped) Base() IObject        { return s.base }
func (s *Scoped) Ctx() context.Context { return s.ctxProvider() }

func (s *Scoped) Err() error       { return s.err }
func (s *Scoped) SetErr(err error) { s.err = err }
