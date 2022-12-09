package goac

import "github.com/heffcodex/goac/oap"

type Embed struct {
	Permissions []oap.Path `json:"__permissions,omitempty"`
}

func NewEmbed(permissions []oap.Path) Embed {
	return Embed{
		Permissions: permissions,
	}
}
