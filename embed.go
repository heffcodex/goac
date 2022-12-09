package goac

type Embed struct {
	Permissions []OAPPath `json:"__permissions,omitempty"`
}

func NewEmbed(permissions []OAPPath) Embed {
	return Embed{
		Permissions: permissions,
	}
}
