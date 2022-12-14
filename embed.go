package goac

import (
	"context"
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	GetPermissions() []goacoap.Path
	SetPermissions([]goacoap.Path)
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []goacoap.Path `json:"__permissions,omitempty"`
}

func (e *Embed) GetPermissions() []goacoap.Path {
	return e.Permissions
}

func (e *Embed) SetPermissions(p []goacoap.Path) {
	e.Permissions = p
}

type ObjectEnforcer[T IEmbed] func(ctx context.Context, o T) (T, error)

func EnforcerFromContext[T IEmbed](ctx context.Context, key string) ObjectEnforcer[T] {
	return ctx.Value(key).(ObjectEnforcer[T])
}

func EnforcerToContext[T IEmbed](ctx context.Context, key string, enforcer ObjectEnforcer[T]) context.Context {
	return context.WithValue(ctx, key, enforcer)
}
