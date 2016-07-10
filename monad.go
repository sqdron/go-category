package category

type Monad interface {
	Return(value interface{}) Monad
	Bind(func(interface{}) Monad) Monad
	GetValue() interface{}
}