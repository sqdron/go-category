package category;

type IApplicative interface {
	IFunctor
	//Pure(interface{}) IFunctor;
	A(f IFunctor) IApplicative;
}