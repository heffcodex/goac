package goacoap

type IParameter interface {
	IParametrized
	SetAllow(v bool) IParameter
	Allow() IParameter
	Deny() IParameter
	End() IParameter
	Finalize() IParameter
}

var _ IParameter = (*Parameter)(nil)

type Parameter struct {
	parent    *Parameter
	name      string
	params    map[string]*Parameter
	allowed   bool
	deadEnd   bool
	finalized bool
}

func (p *Parameter) Name() string               { return p.name }
func (p *Parameter) SetAllow(v bool) IParameter { p.setAllow(v); return p }
func (p *Parameter) Allow() IParameter          { p.setAllow(true); return p }
func (p *Parameter) Deny() IParameter           { p.setAllow(false); return p }
func (p *Parameter) Allowed() bool              { return p.allowed }

func (p *Parameter) Param(name string) IParameter {
	if p.deadEnd {
		panic("cannot append to dead-end node")
	}

	param, ok := p.params[name]
	if !ok {
		if p.finalized {
			panic("cannot append to finalized node")
		}

		param = &Parameter{parent: p, name: name}
		p.params[name] = param
	}

	return param
}

func (p *Parameter) End() IParameter {
	if p.finalized {
		panic("cannot modify finalized node")
	} else if len(p.params) > 0 {
		panic("cannot set node with children as dead-end")
	}

	p.deadEnd = true

	return p
}

func (p *Parameter) Permissions() []Permission {
	if len(p.params) == 0 {
		return []Permission{NewPermission(p.name)}
	}

	perms := make([]Permission, 0, len(p.params))

	for _, param := range p.params {
		if !param.allowed {
			continue
		}

		for _, perm := range param.Permissions() {
			perms = append(perms, NewPermission(p.name).AppendPath(perm.String()))
		}
	}

	return perms
}

func (p *Parameter) Finalize() IParameter {
	p.finalized = true

	for _, param := range p.params {
		param.Finalize()
	}

	return p
}

func (p *Parameter) setAllow(v bool, inPropagation ...bool) {
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
