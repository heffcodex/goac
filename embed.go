package goac

import "github.com/heffcodex/goac/oap"

type Embed struct {
	Permissions []goacoap.Path `json:"__permissions,omitempty"`
}

func NewEmbed(o goacoap.IObject) Embed {
	return Embed{
		Permissions: o.GetAllowedPaths(),
	}
}
