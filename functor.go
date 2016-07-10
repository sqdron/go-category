package category

type Functor interface {
	Fmap(f Morphism, functor Functor) Functor
}