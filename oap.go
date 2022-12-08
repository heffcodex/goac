package goac

type Name string
type Path string

type INode interface {
	Name() Name
	GetAllowedPaths() []Path
}

type IObject interface {
	INode
	Action(name Name) *Action
	Finalize() *Object
}

type IParametrized interface {
	INode
	Allow()
	Deny()
	IsAllowed() bool
	Param(name Name) *Param
	End() *Param
}

type IAction interface {
	IParametrized
	Finalize() *Action
}

type IParam interface {
	IParametrized
	Finalize() *Param
}

var (
	_ IObject = (*Object)(nil)
	_ IAction = (*Action)(nil)
	_ IParam  = (*Param)(nil)
)

type Object struct {
	name      Name
	actions   map[Name]*Action
	finalized bool
}

func NewObject(name Name) *Object {
	return &Object{
		name:    name,
		actions: make(map[Name]*Action),
	}
}

func (o *Object) Name() Name { return o.name }

func (o *Object) Action(name Name) *Action {
	if o.actions == nil {
		o.actions = make(map[Name]*Action)
	}

	action, ok := o.actions[name]
	if !ok {
		if o.finalized {
			panic("cannot append to finalized object")
		}

		action = &Action{
			Param: Param{name: name},
		}
		o.actions[name] = action
	}

	return action
}

func (o *Object) GetAllowedPaths() []Path {
	paths := make([]Path, 0, len(o.actions))

	for _, action := range o.actions {
		if !action.allowed {
			continue
		}

		for _, path := range action.GetAllowedPaths() {
			paths = append(paths, Path(o.name)+"."+path)
		}
	}

	return paths
}

func (o *Object) Finalize() *Object {
	o.finalized = true

	for _, action := range o.actions {
		action.Finalize()
	}

	return o
}

type Action struct {
	Param
}

func (a *Action) Finalize() *Action {
	a.Param.Finalize()
	return a
}

type Param struct {
	parent    *Param
	name      Name
	params    map[Name]*Param
	allowed   bool
	deadEnd   bool
	finalized bool
}

func (p *Param) Name() Name      { return p.name }
func (p *Param) Allow()          { p.setAllow(true) }
func (p *Param) Deny()           { p.setAllow(false) }
func (p *Param) IsAllowed() bool { return p.allowed }

func (p *Param) Param(name Name) *Param {
	if p.deadEnd {
		panic("cannot append to dead-end node")
	}

	param, ok := p.params[name]
	if !ok {
		if p.finalized {
			panic("cannot append to finalized node")
		}

		param = &Param{parent: p, name: name}
		p.params[name] = param
	}

	return param
}

func (p *Param) End() *Param {
	if p.finalized {
		panic("cannot modify finalized node")
	} else if len(p.params) > 0 {
		panic("cannot set node with children as dead-end")
	}

	p.deadEnd = true

	return p
}

func (p *Param) GetAllowedPaths() []Path {
	paths := make([]Path, 0, len(p.params))

	for _, param := range p.params {
		if !param.allowed {
			continue
		}

		for _, path := range param.GetAllowedPaths() {
			paths = append(paths, Path(p.name)+"."+path)
		}
	}

	return paths
}

func (p *Param) Finalize() *Param {
	p.finalized = true

	for _, param := range p.params {
		param.Finalize()
	}

	return p
}

func (p *Param) setAllow(v bool, inPropagation ...bool) {
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
