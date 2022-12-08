package goac

type Embed struct {
	Permissions []Path `json:"__permissions,omitempty"`
}

func (e *Embed) SetOAPPermissions(o IObject) {
	e.Permissions = o.GetAllowedPaths()
}
