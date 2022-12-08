package goac

import "encoding/json"

type IEmbed interface {
	json.Marshaler
	GetOAP() IOAPObject
	SetOAP(o IOAPObject)
	UpdateOAPPermissions()
}

var _ IEmbed = (*Embed)(nil)

type Embed struct {
	Permissions []OAPPath `json:"__permissions,omitempty"`
	o           IOAPObject
}

func (e *Embed) GetOAP() IOAPObject {
	return e.o
}

func (e *Embed) SetOAP(o IOAPObject) {
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
