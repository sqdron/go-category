package theory;

type IMaybe interface {
	IApplicative
}

type just struct {
	value interface{}
};

type nothing struct {
};

func Maybe(val interface{}) Category {
	if (val == nil){
		return Nothing();
	};
	return Just(val);
}

func Just(value interface{}) Category {
	var res Category = Category{};
	res.context = just{value : value};
	return res;
}

func Nothing() Category {
	return Category{context:nothing{}}
}

// fmap f (Just val) = Just (f val)
func (val just) FMap(f Morphism) Category{
	return Maybe(f(val.value));
}

// fmap f Nothing = Nothing
func (val nothing) FMap(f Morphism) Category {
	return Nothing();
}

func (val just) AMap(jm Category) Category{
	var appl = jm.Value();
	if (appl == empty{}) {
		return Nothing()
	}
	var f Morphism = appl.(Morphism);
	return Maybe(f(val.value));
}

// fmap f Nothing = Nothing
func (val nothing) AMap(jm Category) Category {
	return Nothing();
}

func (val just) Value() interface {}{
	return val.value;
}

func (val nothing) Value() interface{} {
	return empty{};
}
func (mp fMap) Just(val interface{}) Category {
	return mp(Just(val));
}

func (mp fMap) Nothing() Category {
	return Nothing();
}