package theory;

//type Functor func(interface{})interface{};

type IFunctor interface {
	Fmap(f Morphism) IFunctor;
}

//func (m Morphism) S(f IFunctor) IFunctor{
//	return f.Fmap(m);
//}