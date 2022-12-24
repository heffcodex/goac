package goac

import (
	"context"
	"fmt"
	goacoap "github.com/heffcodex/goac/oap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFromContext(t *testing.T) {
	type O struct {
		Embed
	}

	var enf Enforcer[*O] = func(ctx context.Context, obj *O, _ goacoap.IObject) (*O, error) {
		return obj, fmt.Errorf("test")
	}

	assert.PanicsWithValue(t, "context is nil", func() {
		FromContext[*O](nil)
	})
	assert.PanicsWithValue(t, "enforcer not found in context", func() {
		FromContext[*O](context.Background())
	})
	assert.NotPanics(t, func() {
		FromContext[*O](nil, enf)
	})
	assert.NotPanics(t, func() {
		FromContext[*O](context.Background(), enf)
	})

	ctx := context.WithValue(context.Background(), ctxKey[*O]{}, enf)

	_, err1 := FromContext[*O](ctx)(nil, nil, nil)
	_, err2 := enf(nil, nil, nil)
	assert.EqualError(t, err2, err1.Error())
}

func TestToContext(t *testing.T) {
	type O struct {
		Embed
	}

	var enf Enforcer[*O] = func(ctx context.Context, obj *O, _ goacoap.IObject) (*O, error) {
		return obj, fmt.Errorf("test")
	}

	ctx := context.Background()
	ctx = ToContext[*O](ctx, enf)

	_, err1 := FromContext[*O](ctx)(nil, nil, nil)
	_, err2 := enf(nil, nil, nil)

	assert.EqualError(t, err2, err1.Error())
}

func TestEnforce(t *testing.T) {
	type O struct {
		Embed
		I int
	}

	type OAP struct {
		goacoap.Object
	}

	object := &O{I: 1}

	oapObject := &OAP{}
	goacoap.Build(oapObject, "oap")
	oapObject.Allow(goacoap.True())

	var enf Enforcer[*O] = func(ctx context.Context, obj *O, oapObj goacoap.IObject) (*O, error) {
		obj.I++
		return obj, nil
	}

	ctx := context.Background()
	ctx = ToContext[*O](ctx, enf)

	o, err := Enforce[*O](ctx, object, oapObject)
	require.NoError(t, err)
	assert.Equal(t, 2, o.I)
	assert.EqualValues(t, []goacoap.Permission{"oap"}, o.Embed.Permissions)
}
