package category

//type Functor func(interface{})interface{};

type IFunctor interface {
	Fmap(f Morphism) IFunctor;
}

