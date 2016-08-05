package theory

type Morphism func(interface{}) interface{}

type Category struct {
	context interface{}
}

type IFunctor interface {
	Map(Morphism) Category
	Value() interface{}
}

type IApplicative interface {
	IFunctor
	Applicative(jm Category) Category
}

type IMonad interface {
	IApplicative
	Bind(func(interface{}) Category) Category
}

type fMap func(j Category) Category

func (cat Category) FMap(m Morphism) Category {
	return cat.context.(IFunctor).Map(m)
}

func (cat Category) Value() interface{} {
	return cat.context.(IFunctor).Value()
}

func (cat Category) Bind(m func(interface{}) Category) Category {
	return cat.context.(IMonad).Bind(m)
}

func (m Morphism) S() fMap {
	return func(c Category) Category {
		return c.context.(IFunctor).Map(m)
	}
}

func (jm Category) A() fMap {
	return func(c Category) Category {
		return c.context.(IApplicative).Applicative(jm)
	}
}
