package goac

import "context"

type Enforced[T IEmbed] *T
type Enforcer[T IEmbed] func(ctx context.Context, object T) (Enforced[T], error)

func EnforcerFromContext[T IEmbed](ctx context.Context, key string) Enforcer[T] {
	return ctx.Value(key).(Enforcer[T])
}

func EnforcerToContext[T IEmbed](ctx context.Context, key string, enforcer Enforcer[T]) context.Context {
	return context.WithValue(ctx, key, enforcer)
}
