package goac

import "context"

type Enforcer[T IEmbed] func(ctx context.Context, object T) (T, error)

func EnforcerFromContext[T IEmbed](ctx context.Context, key string) Enforcer[T] {
	return ctx.Value(key).(Enforcer[T])
}

func EnforcerFromContextWithFallback[T IEmbed](ctx context.Context, key string, fallback ...Enforcer[T]) Enforcer[T] {
	_fallback := func(_ context.Context, object T) (T, error) {
		return object, nil
	}

	if len(fallback) > 1 {
		panic("fallback must be a single function")
	} else if len(fallback) == 1 && fallback[0] != nil {
		_fallback = fallback[0]
	}

	if ctx == nil {
		return _fallback
	}

	enforcer, ok := ctx.Value(key).(Enforcer[T])
	if !ok {
		return _fallback
	}

	return enforcer
}

func EnforcerToContext[T IEmbed](ctx context.Context, key string, enforcer Enforcer[T]) context.Context {
	return context.WithValue(ctx, key, enforcer)
}
