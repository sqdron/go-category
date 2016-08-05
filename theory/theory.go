package theory

//Morphism is function defenition type
type Morphism func(interface{}) interface{}

//Category base category struct
type Category struct {
	context interface{}
}

//IFunctor is a functor interface
type IFunctor interface {
	Map(Morphism) Category
	Value() interface{}
}

//IApplicative is an applicative interface
type IApplicative interface {
	IFunctor
	Applicative(jm Category) Category
}

//IMonad is a Monad interface
type IMonad interface {
	IApplicative
	Bind(func(interface{}) Category) Category
}

type fMap func(j Category) Category

//FMap is a Haskel fmap implementation
func (cat Category) FMap(m Morphism) Category {
	return cat.context.(IFunctor).Map(m)
}

//Value extracts from a functor
func (cat Category) Value() interface{} {
	return cat.context.(IFunctor).Value()
}

//Bind haskel implementation
func (cat Category) Bind(m func(interface{}) Category) Category {
	return cat.context.(IMonad).Bind(m)
}

//S is a $ in a Haskel style implementation
func (m Morphism) S() fMap {
	return func(c Category) Category {
		return c.context.(IFunctor).Map(m)
	}
}

//A is a * in a Haskel style implementation
func (cat Category) A() fMap {
	return func(c Category) Category {
		return c.context.(IApplicative).Applicative(cat)
	}
}
