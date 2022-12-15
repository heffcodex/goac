package goacoap

type IObject interface {
	INode
	Action(name string) IAction
	Fresh() IObject
	Finalize() IObject
}

var _ IObject = (*Object)(nil)

type Object struct {
	name      string
	actions   map[string]*Action
	finalized bool
}

func NewObject(name string) *Object {
	return &Object{
		name:    name,
		actions: make(map[string]*Action),
	}
}

func (o *Object) Name() string { return o.name }

func (o *Object) Action(name string) IAction {
	if o.actions == nil {
		o.actions = make(map[string]*Action)
	}

	action, ok := o.actions[name]
	if !ok {
		if o.finalized {
			panic("cannot append to finalized object")
		}

		action = &Action{
			Parameter: Parameter{name: name},
		}
		o.actions[name] = action
	}

	return action
}

func (o *Object) Permissions() []Permission {
	perms := make([]Permission, 0, len(o.actions))

	for _, action := range o.actions {
		if !action.allowed {
			continue
		}

		for _, perm := range action.Permissions() {
			perms = append(perms, NewPermission(o.name).AppendPath(perm.String()))
		}
	}

	return perms
}

func (o *Object) Fresh() IObject {
	return NewObject(o.name)
}

func (o *Object) Finalize() IObject {
	o.finalized = true

	for _, action := range o.actions {
		action.Finalize()
	}

	return o
}
