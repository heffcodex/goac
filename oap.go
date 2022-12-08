package goac

type OAPName string
type OAPPath string

type IOAPNode interface {
	Name() OAPName
	GetAllowedPaths() []OAPPath
}

type IOAPObject interface {
	IOAPNode
	Action(name OAPName) *OAPAction
	Finalize() *OAPObject
}

type IOAPParametrized interface {
	IOAPNode
	Allow()
	Deny()
	IsAllowed() bool
	Param(name OAPName) *OAPParam
	End() *OAPParam
}

type IOAPAction interface {
	IOAPParametrized
	Finalize() *OAPAction
}

type IOAPParam interface {
	IOAPParametrized
	Finalize() *OAPParam
}

var (
	_ IOAPObject = (*OAPObject)(nil)
	_ IOAPAction = (*OAPAction)(nil)
	_ IOAPParam  = (*OAPParam)(nil)
)

type OAPObject struct {
	name      OAPName
	actions   map[OAPName]*OAPAction
	finalized bool
}

func NewOAPObject(name OAPName) *OAPObject {
	return &OAPObject{
		name:    name,
		actions: make(map[OAPName]*OAPAction),
	}
}

func (o *OAPObject) Name() OAPName { return o.name }

func (o *OAPObject) Action(name OAPName) *OAPAction {
	if o.actions == nil {
		o.actions = make(map[OAPName]*OAPAction)
	}

	action, ok := o.actions[name]
	if !ok {
		if o.finalized {
			panic("cannot append to finalized object")
		}

		action = &OAPAction{
			OAPParam: OAPParam{name: name},
		}
		o.actions[name] = action
	}

	return action
}

func (o *OAPObject) GetAllowedPaths() []OAPPath {
	paths := make([]OAPPath, 0, len(o.actions))

	for _, action := range o.actions {
		if !action.allowed {
			continue
		}

		for _, path := range action.GetAllowedPaths() {
			paths = append(paths, OAPPath(o.name)+"."+path)
		}
	}

	return paths
}

func (o *OAPObject) Finalize() *OAPObject {
	o.finalized = true

	for _, action := range o.actions {
		action.Finalize()
	}

	return o
}

type OAPAction struct {
	OAPParam
}

func (a *OAPAction) Finalize() *OAPAction {
	a.OAPParam.Finalize()
	return a
}

type OAPParam struct {
	parent    *OAPParam
	name      OAPName
	params    map[OAPName]*OAPParam
	allowed   bool
	deadEnd   bool
	finalized bool
}

func (p *OAPParam) Name() OAPName   { return p.name }
func (p *OAPParam) Allow()          { p.setAllow(true) }
func (p *OAPParam) Deny()           { p.setAllow(false) }
func (p *OAPParam) IsAllowed() bool { return p.allowed }

func (p *OAPParam) Param(name OAPName) *OAPParam {
	if p.deadEnd {
		panic("cannot append to dead-end node")
	}

	param, ok := p.params[name]
	if !ok {
		if p.finalized {
			panic("cannot append to finalized node")
		}

		param = &OAPParam{parent: p, name: name}
		p.params[name] = param
	}

	return param
}

func (p *OAPParam) End() *OAPParam {
	if p.finalized {
		panic("cannot modify finalized node")
	} else if len(p.params) > 0 {
		panic("cannot set node with children as dead-end")
	}

	p.deadEnd = true

	return p
}

func (p *OAPParam) GetAllowedPaths() []OAPPath {
	paths := make([]OAPPath, 0, len(p.params))

	for _, param := range p.params {
		if !param.allowed {
			continue
		}

		for _, path := range param.GetAllowedPaths() {
			paths = append(paths, OAPPath(p.name)+"."+path)
		}
	}

	return paths
}

func (p *OAPParam) Finalize() *OAPParam {
	p.finalized = true

	for _, param := range p.params {
		param.Finalize()
	}

	return p
}

func (p *OAPParam) setAllow(v bool, inPropagation ...bool) {
	if p.finalized {
		panic("cannot modify finalized node")
	}

	_inPropagation := len(inPropagation) > 0 && inPropagation[0]

	if v {
		if !p.deadEnd && !_inPropagation {
			panic("cannot allow non-dead-end node")
		}

		p.allowed = true

		if p.parent != nil {
			p.parent.setAllow(true, true)
		}
	} else {
		p.allowed = false

		for _, child := range p.params {
			child.setAllow(false, true)
		}
	}
}
