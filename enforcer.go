package goac

import (
	"context"
	goacoap "github.com/heffcodex/goac/oap"
)

type ctxKey[T IEmbed] struct{ _ T }

type Enforcer[T IEmbed] func(ctx context.Context, obj T, oapObj goacoap.IObject) (T, error)

func FromContext[T IEmbed](ctx context.Context, fallback ...Enforcer[T]) Enforcer[T] {
	var _fallback Enforcer[T]

	if len(fallback) > 1 {
		panic("fallback must be a single function")
	} else if len(fallback) == 1 && fallback[0] != nil {
		_fallback = fallback[0]
	}

	if ctx == nil {
		if _fallback != nil {
			return _fallback
		}

		panic("context is nil")
	}

	enforcer, ok := ctx.Value(ctxKey[T]{}).(Enforcer[T])
	if !ok {
		if _fallback != nil {
			return _fallback
		}

		panic("enforcer not found in context")
	}

	return enforcer
}

func ToContext[T IEmbed](ctx context.Context, enforcer Enforcer[T]) context.Context {
	return context.WithValue(ctx, ctxKey[T]{}, enforcer)
}

func Enforce[T IEmbed](ctx context.Context, obj T, oapObj goacoap.IObject, fallback ...Enforcer[T]) (T, error) {
	enforce := FromContext[T](ctx, fallback...)
	enforced, err := enforce(ctx, obj, oapObj)
	if err != nil {
		return *new(T), err
	}

	allowed := goacoap.CollectAllowed(oapObj)
	obj.setPermissions(allowed)

	return enforced, nil
}
