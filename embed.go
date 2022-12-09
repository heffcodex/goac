package goac

import "github.com/heffcodex/goac/oap"

type Embed struct {
	Permissions []oap.Path `json:"__permissions,omitempty"`
}

func NewEmbed(o oap.IObject) Embed {
	return Embed{
		Permissions: o.GetAllowedPaths(),
	}
}
