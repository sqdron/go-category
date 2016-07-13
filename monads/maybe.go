package monads
//
//type Just struct {
//	value interface{}
//}
//
//type Nothing struct {}
//
//type IMaybe interface {
//	Pure(a interface{}) IMaybe;
//	Bind(fn func(a interface{}) IMaybe) IMaybe;
//}
//
//func Maybe(val interface{}) IMaybe{
//	if (val == nil){
//		return Nothing{};
//	}
//	return Just{value:val}
//}
//
//func (a Just) Pure(a interface{}) IMaybe {
//	return Just{value:a};
//}
//
//func (a Nothing) Pure(a interface{}) IMaybe {
//	return Just{value:a};
//}
