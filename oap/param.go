package oap

type IParam interface {
	IParametrized
	SetAllow(v bool) *Param
	Allow() *Param
	Deny() *Param
	End() *Param
	Finalize() *Param
}

var _ IParam = (*Param)(nil)

type Param struct {
	parent    *Param
	name      string
	params    map[string]*Param
	allowed   bool
	deadEnd   bool
	finalized bool
}

func (p *Param) Name() string           { return p.name }
func (p *Param) SetAllow(v bool) *Param { p.setAllow(v); return p }
func (p *Param) Allow() *Param          { p.setAllow(true); return p }
func (p *Param) Deny() *Param           { p.setAllow(false); return p }
func (p *Param) Allowed() bool          { return p.allowed }

func (p *Param) Param(name string) *Param {
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
	if len(p.params) == 0 {
		return []Path{NewPath(p.name)}
	}

	paths := make([]Path, 0, len(p.params))

	for _, param := range p.params {
		if !param.allowed {
			continue
		}

		for _, path := range param.GetAllowedPaths() {
			paths = append(paths, NewPath(p.name).Append(path.String()))
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
