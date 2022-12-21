package goac

import (
	"context"
	"errors"
	gojson "github.com/goccy/go-json"
	goacoap "github.com/heffcodex/goac/oap"
)

type Enforcer[T any] func(ctx context.Context, object T) (*Enforced[T], error)

func EnforcerFromContext[T any](ctx context.Context, key string) Enforcer[T] {
	return ctx.Value(key).(Enforcer[T])
}

func EnforcerToContext[T any](ctx context.Context, key string, enforcer Enforcer[T]) context.Context {
	return context.WithValue(ctx, key, enforcer)
}

var _ gojson.Marshaler = (*Enforced)(nil)

type Enforced[T any] struct {
	object T
	perms  []goacoap.Permission
}

func NewEnforced[T any](object T, perms []goacoap.Permission) *Enforced[T] {
	return &Enforced[T]{
		object: object,
		perms:  perms,
	}
}

func (e *Enforced[T]) Permissions(perms ...[]goacoap.Permission) []goacoap.Permission {
	if len(perms) > 1 {
		panic("too many arguments")
	} else if len(perms) == 1 {
		e.setPermissions(perms[0])
	}

	return e.getPermissions()
}

func (e *Enforced[T]) getPermissions() []goacoap.Permission {
	perms := make([]goacoap.Permission, len(e.perms))
	copy(perms, e.perms)
	return perms
}

func (e *Enforced[T]) setPermissions(perms []goacoap.Permission) {
	e.perms = make([]goacoap.Permission, len(perms))
	copy(e.perms, perms)
}

func (e Enforced[T]) MarshalJSON() ([]byte, error) {
	objectBytes, err := gojson.Marshal(e.object)
	if err != nil {
		return nil, errors.New("marshal object")
	}

	if len(objectBytes) == 0 {
		return objectBytes, nil
	} else if objectBytes[0] != '{' {
		return nil, errors.New("enforced object must be a JSON object")
	}

	permsBytes, err := gojson.Marshal(e.perms)
	if err != nil {
		return nil, errors.New("marshal permissions")
	}

	bytes := append(
		objectBytes[:1],
		append(
			append(
				[]byte(`"__permissions":`),
				permsBytes...,
			),
			objectBytes[1:]...,
		)...,
	)

	return bytes, nil
}
