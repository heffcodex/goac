package goac

import (
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	GetPermissions() []goacoap.Permission
	SetPermissions(perms []goacoap.Permission)
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []goacoap.Permission `json:"__permissions,omitempty"`
}

func (e *Embed) GetPermissions() []goacoap.Permission {
	perms := make([]goacoap.Permission, len(e.Permissions))
	copy(perms, e.Permissions)
	return perms
}

func (e *Embed) SetPermissions(perms []goacoap.Permission) {
	e.Permissions = make([]goacoap.Permission, len(perms))
	copy(e.Permissions, perms)
}
