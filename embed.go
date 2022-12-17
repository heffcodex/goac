package goac

import (
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	Permissions(perms ...[]goacoap.Permission) []goacoap.Permission
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	P []goacoap.Permission `json:"__permissions,omitempty"`
}

func (e *Embed) Permissions(perms ...[]goacoap.Permission) []goacoap.Permission {
	if len(perms) > 1 {
		panic("too many arguments")
	} else if len(perms) == 1 {
		e.P = make([]goacoap.Permission, len(perms[0]))
		copy(e.P, perms[0])
	}

	read := make([]goacoap.Permission, len(e.P))
	copy(read, e.P)
	return read
}
