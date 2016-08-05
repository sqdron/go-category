package theory

//IMaybe is na Maybe monad implementation interface
type IMaybe interface {
	IMonad
}

type just struct {
	value interface{}
}

type nothing struct {
}

//Maybe constructor
func Maybe(val interface{}) Category {
	if val == nil {
		return Nothing()
	}
	return Just(val)
}

//Just constructor
func Just(value interface{}) Category {
	return Category{context: just{value: value}}
}

//Nothing constructor
func Nothing() Category {
	return Category{context: nothing{}}
}

//Map - fmap f (Just val) = Just (f val)
func (val just) Map(f Morphism) Category {
	return Maybe(f(val.value))
}

//Map fmap f Nothing = Nothing
func (val nothing) Map(f Morphism) Category {
	return Nothing()
}

//Applicative implementation
func (val just) Applicative(jm Category) Category {
	var appl = jm.Value()
	if (appl == empty{}) {
		return Nothing()
	}
	return Maybe(appl.(Morphism)(val.value))
}

//Applicative fmap f Nothing = Nothing
func (val nothing) Applicative(jm Category) Category {
	return Nothing()
}

//Bind implementation
func (val just) Bind(f func(interface{}) Category) Category {
	return f(val.value)
}

//Bind implementation
func (val nothing) Bind(func(interface{}) Category) Category {
	return Nothing()
}

//Value implementation
func (val just) Value() interface{} {
	return val.value
}

//Value implementation
func (val nothing) Value() interface{} {
	return empty{}
}

//Just implementation
func (mp fMap) Just(val interface{}) Category {
	return mp(Just(val))
}

//Nothing implementation
func (mp fMap) Nothing() Category {
	return mp(Nothing())
}
