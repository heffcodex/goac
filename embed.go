package goac

type Embed struct {
	Permissions []OAPPath `json:"__permissions,omitempty"`
}

func NewEmbed(oap IOAPObject) Embed {
	return Embed{
		Permissions: oap.GetAllowedPaths(),
	}
}
