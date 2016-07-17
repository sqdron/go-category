package theory;

type Morphism func(interface{}) interface{};

type Category struct {
	context interface{};
}

type IFunctor interface {
	FMap(Morphism) Category;
	Value() interface{};
}

type IApplicative interface {
	IFunctor
	AMap(jm Category) Category;
}

type fMap func(j Category) Category;

func (cat Category) FMap(m Morphism) Category {
	return cat.context.(IFunctor).FMap(m);
}

func (cat Category) Value() interface{} {
	return cat.context.(IFunctor).Value();
}

func (m Morphism) S() fMap {
	return func(c Category) Category{
		return c.context.(IFunctor).FMap(m);
	}
}

func (jm Category) A() fMap {
	return func(c Category) Category{
		return c.context.(IApplicative).AMap(jm);
	}
}