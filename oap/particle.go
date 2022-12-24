package oap

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

var (
	_          IParticle = (*particle)(nil)
	iParticleT           = reflect.TypeOf((*IParticle)(nil)).Elem()
)

type AllowFn func(p IParticle) bool

func True() AllowFn         { return Static(true) }
func False() AllowFn        { return Static(false) }
func Static(v bool) AllowFn { return func(IParticle) bool { return v } }

type private struct{}

type IParticle interface {
	Name() string
	Allowed() bool
	Allow(f AllowFn)

	setName(name string)
	isBuilt() bool
	setBuilt()
	parent() IParticle
	setParent(IParticle)
	children() []IParticle
	setChildren(children []IParticle)
	path() Permission

	__(private)
}

type particle struct {
	_parent   IParticle
	_children []IParticle
	name      string
	allow     AllowFn
	built     bool
}

func (p *particle) Name() string { return p.name }

func (p *particle) Allowed() bool {
	mustBuiltEnd(p)

	if p.allow == nil {
		return false
	}

	return p.allow(p)
}

func (p *particle) Allow(f AllowFn) {
	mustBuiltEnd(p)
	p.allow = f
}

func (p *particle) setName(name string)              { p.name = name }
func (p *particle) isBuilt() bool                    { return p.built }
func (p *particle) setBuilt()                        { p.built = true }
func (p *particle) parent() IParticle                { return p._parent }
func (p *particle) setParent(parent IParticle)       { p._parent = parent }
func (p *particle) children() []IParticle            { return p._children }
func (p *particle) setChildren(children []IParticle) { p._children = children }
func (p *particle) path() Permission {
	thisPerm := NewPermission(p.Name())
	if p.parent() == nil {
		return thisPerm
	}

	return p.parent().path().AppendPath(thisPerm)
}

func (p *particle) __(private) {}

func mustBuiltEnd(p IParticle) {
	if !p.isBuilt() {
		panic("not built")
	} else if len(p.children()) > 0 {
		panic("not end of tree")
	}
}

func build(this, parent IParticle) {
	vof := reflect.ValueOf(this).Elem()
	tof := vof.Type()

	if tof.Kind() != reflect.Struct {
		panic("not a struct")
	}

	if name := this.Name(); name == "" {
		this.setName(strcase.ToSnake(tof.Name()))
	}

	children := make([]IParticle, 0, vof.NumField())

	for i := 0; i < vof.NumField(); i++ {
		v := vof.Field(i)
		t := tof.Field(i)

		if v.Kind() == reflect.Struct && !t.Anonymous && t.IsExported() && v.Addr().Type().Implements(iParticleT) {
			child := v.Addr().Interface().(IParticle)
			child.setName(strcase.ToSnake(t.Name))

			build(child, this)
			children = append(children, child)
		}
	}

	this.setParent(parent)
	this.setChildren(children)
	this.setBuilt()
}

func collectPermissions(p IParticle, checkFn func(p IParticle) bool, acc []Permission) []Permission {
	if len(p.children()) == 0 && checkFn(p) {
		acc = append(acc, p.path())
		return acc
	}

	for _, child := range p.children() {
		acc = collectPermissions(child, checkFn, acc)
	}

	return acc
}
