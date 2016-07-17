package theory;

type right func() any;

type left func() any

type IEither interface {
	IMonad
}