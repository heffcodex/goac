package goac

import (
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	setPermissions(perms []oap.Permission)
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []oap.Permission `json:"__permissions,omitempty"`
}

func (e *Embed) setPermissions(perms []oap.Permission) {
	e.Permissions = perms
}
