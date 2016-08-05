package theory

type right func() any

type left func() any

//IEither is na Either monad implementation
type IEither interface {
	IMonad
}
