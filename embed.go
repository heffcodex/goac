package goac

import (
	"github.com/heffcodex/goac/oap"
)

type IEmbed interface {
	GetPermissions() []goacoap.Path
	SetPermissions([]goacoap.Path)
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []goacoap.Path `json:"__permissions,omitempty"`
}

func (e *Embed) GetPermissions() []goacoap.Path {
	return e.Permissions
}

func (e *Embed) SetPermissions(p []goacoap.Path) {
	e.Permissions = p
}
