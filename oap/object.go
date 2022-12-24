package oap

var (
	_ IObject    = (*Object)(nil)
	_ IAction    = (*Action)(nil)
	_ IParameter = (*Parameter)(nil)
)

type IObject interface{ IParticle }
type Object struct{ particle }

func (o *Object) setParent(IParticle) {}

type IAction interface{ IParameter }
type Action struct{ Parameter }

type IParameter interface{ IParticle }
type Parameter struct{ particle }

func Build(o IObject, name ...string) {
	if len(name) > 1 {
		panic("too many arguments")
	} else if len(name) == 1 {
		o.setName(name[0])
	}

	build(o, nil)
}

func CollectAll(o IObject) []Permission {
	return collectPermissions(o, func(p IParticle) bool { return true }, []Permission{})
}

func CollectAllowed(o IObject) []Permission {
	return collectPermissions(o, func(p IParticle) bool { return p.Allowed() }, []Permission{})
}
