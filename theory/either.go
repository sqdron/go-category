package theory;

type right func() any;

type left func() any

type IEither interface {
	IMonad
}

//
//func Left(value interface{}) Category {
//	var res Category = Category{};
//	res.context = just{value : value};
//	return res;
//}
//
//func Right() Category {
//	return Category{context:nothing{}}
//}
//
//// fmap f (Just val) = Just (f val)
//func (val left) FMap(f Morphism) Category{
//	return Maybe(f(val.value));
//}
//
//// fmap f Nothing = Nothing
//func (val right) FMap(f Morphism) Category {
//	return Nothing();
//}
//
//func (val just) AMap(jm Category) Category{
//	var appl = jm.Value();
//	if (appl == empty{}) {
//		return Nothing()
//	}
//	var f Morphism = appl.(Morphism);
//	return Maybe(f(val.value));
//}
//
//// fmap f Nothing = Nothing
//func (val nothing) AMap(jm Category) Category {
//	return Nothing();
//}
//
//func (val just) Bind(f func(interface{}) Category) Category{
//	return f(val.value);
//}
//
//func (val nothing) Bind(func(interface{}) Category) Category{
//	return Nothing();
//}
//
//func (val just) Value() interface {}{
//	return val.value;
//}
//
//func (val nothing) Value() interface{} {
//	return empty{};
//}
//func (mp fMap) Just(val interface{}) Category {
//	return mp(Just(val));
//}
//
//func (mp fMap) Nothing() Category {
//	return mp(Nothing());
//}
