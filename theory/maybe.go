package theory;

type IMaybe interface {
	IApplicative
}

func Maybe(val interface{}) IMaybe {
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
	return ft.Fmap(applicativeFunc).(IMaybe);
}


type FMapFunc func(IFunctor) IFunctor

func (m FMapFunc) Just(value interface{}) IFunctor {
	return m(J(value));
}

func (m FMapFunc) Nothing(value interface{}) IFunctor {
	return N();
}

func (m Morphism) S() FMapFunc {
	return func(j IFunctor) IFunctor{
		return j.Fmap(m);
	}
}