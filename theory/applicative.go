package theory;

type IApplicative interface {
	IFunctor
	A(f IFunctor) IApplicative;
}

//func (m Morphism) S(f IFunctor) IApplicative {
//	return f.Fmap(m).(IApplicative);
//}

