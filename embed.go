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
		e.setPermissions(perms[0])
	}

	return e.getPermissions()
}

func (e *Embed) getPermissions() []goacoap.Permission {
	perms := make([]goacoap.Permission, len(e.P))
	copy(perms, e.P)
	return perms
}

func (e *Embed) setPermissions(perms []goacoap.Permission) {
	e.P = make([]goacoap.Permission, len(perms))
	copy(e.P, perms)
}
