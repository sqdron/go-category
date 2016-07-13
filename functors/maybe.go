package functors;

import (
	. "github.com/sqdron/go-category"
)

type fMaybe interface {
	IApplicative
}

//type aMaybe interface {
//	IApplicative
//}


func Maybe(val interface{}) fMaybe{
	if (val == nil){
		return N();
	}
	return J(val);
}

func (n Nothing) Fmap(m Morphism) IFunctor{
	return N();
}

func (n Just) Fmap(fn Morphism) IFunctor{
	return Maybe(fn(n.value));
}


func (n Nothing) A(ft IFunctor) IApplicative{
	return N();
}

func (a Just) A(ft IFunctor) IApplicative{
	var applicativeFunc Morphism = a.value.(Morphism);

	return ft.Fmap(applicativeFunc).(fMaybe);
}

