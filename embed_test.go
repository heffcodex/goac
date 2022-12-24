package goac

import (
	"encoding/json"
	goacoap "github.com/heffcodex/goac/oap"
	"github.com/stretchr/testify/require"
	"testing"
)

type TestEmbed struct {
	Embed
	TestValue any
}

func TestEmbed_setPermissions(t *testing.T) {
	s := &TestEmbed{}

	s.setPermissions([]goacoap.Permission{"a", "b", "c"})
	require.Equal(t, []goacoap.Permission{"a", "b", "c"}, s.Permissions)
}

func TestEmbed_MarshalJSON(t *testing.T) {
	s := &TestEmbed{
		TestValue: "test",
	}

	s.setPermissions([]goacoap.Permission{"a", "b", "c"})

	res, err := json.Marshal(s)
	require.NoError(t, err)
	require.Equal(t, `{"__permissions":["a","b","c"],"TestValue":"test"}`, string(res))
}
