package theory;

type right func() any;

type left func() any
//
//func Right(val any) right {
//	return func() any {
//		return val;
//	}
//}
//
//func Left(val interface{}) left {
//	return func() any {
//		return val;
//	}
//}
//
//func (r right) FMap(f Morphism) IFunctor {
//	return Just(f(r()))
//}
//
//func (l left) FMap(f Morphism) IFunctor {
//	return Left(l());
//}