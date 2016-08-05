package theory

type list []interface{}

type IList interface {
	IMonad
}

func List(slice ...interface{}) Category {
	if slice == nil {

		return Category{context: list{}}
	}
	var lst list = slice
	return Category{context: lst}
}

func FromList(lst list) Category {
	return Category{context: lst}
}

func (val list) Map(f Morphism) Category {
	var result list = []interface{}{}
	for _, v := range val {
		result = append(result, f(v))
	}
	return FromList(result)
}

func (val list) Applicative(jm Category) Category {
	var appl = jm.Value()
	if (appl == empty{}) {
		return Nothing()
	}
	var f Morphism = appl.(Morphism)
	var result list = []interface{}{}
	for _, v := range val {
		result = append(result, f(v))
	}
	return FromList(result)
}

// fmap f Nothing = Nothing
func (val list) Bind(f func(interface{}) Category) Category {
	return f(val)
}

func (val list) Value() interface{} {
	return val
}
