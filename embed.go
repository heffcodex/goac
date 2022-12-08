package goac

import "encoding/json"

type IEmbed interface {
	json.Marshaler
	GetOAP() IObject
	SetOAP(o IObject)
	UpdateOAPPermissions()
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []Path `json:"__permissions,omitempty"`
	o           IObject
}

func (e *Embed) GetOAP() IObject {
	return e.o
}

func (e *Embed) SetOAP(o IObject) {
	e.o = o
}

func (e *Embed) UpdateOAPPermissions() {
	if e.o == nil {
		panic("oap is not set")
	}

	e.Permissions = e.o.GetAllowedPaths()
}

func (e *Embed) MarshalJSON() ([]byte, error) {
	e.UpdateOAPPermissions()

	if len(e.Permissions) == 0 {
		return []byte("null"), nil
	}

	return json.Marshal(e.Permissions)
}
