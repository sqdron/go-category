package theory;

type Monad interface {
	Return(value interface{}) Monad
	Bind(func(interface{}, Monad) Monad) Monad
	GetValue() interface{}
}