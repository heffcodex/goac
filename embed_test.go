package goac

import (
	"encoding/json"
	goacoap "github.com/heffcodex/goac/oap"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestEmbed struct {
	*Embed
	TestValue any
}

func TestNewEmbed(t *testing.T) {
	s := &TestEmbed{Embed: NewEmbed(nil)}
	require.Empty(t, s.Permissions())

	obj := goacoap.NewObject("test")
	obj.Action("a").End().Allow()
	obj.Action("b").End().Allow()
	obj.Action("c").End().Allow()

	s = &TestEmbed{Embed: NewEmbed(obj.Permissions())}
	require.Equal(t, []goacoap.Permission{"test.a", "test.b", "test.c"}, s.Permissions())
}

func TestEmbed_MarshalJSON(t *testing.T) {
	s := &TestEmbed{
		Embed: &Embed{
			P: []goacoap.Permission{"a", "b", "c"},
		},
		TestValue: "test",
	}

	res, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, `{"__permissions":["a","b","c"],"TestValue":"test"}`, string(res))
}

func TestEmbed_Permissions(t *testing.T) {
	s := &TestEmbed{
		Embed: &Embed{
			P: []goacoap.Permission{"a", "b", "c"},
		},
	}

	require.Equal(t, []goacoap.Permission{"a", "b", "c"}, s.Permissions())
	require.Equal(t, []goacoap.Permission{"d", "e", "f"}, s.Permissions([]goacoap.Permission{"d", "e", "f"}))
}
