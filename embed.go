package goac

import (
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	Permissions() []goacoap.Permission
}

var _ IEmbed = Embed{}

type Embed struct {
	P []goacoap.Permission `json:"__permissions,omitempty"`
}

func (e Embed) Permissions() []goacoap.Permission {
	perms := make([]goacoap.Permission, len(e.P))
	copy(perms, e.P)
	return perms
}
