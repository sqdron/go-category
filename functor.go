package category

type Functor func(interface{})interface{};

type FunctorClass interface {
	Fmap(f Morphism) func(Functor) Functor;
}

