package theory

type list []interface{}

//IList is an interface for a List Monad
type IList interface {
	IMonad
}

//List constructor
func List(slice ...interface{}) Category {
	if slice == nil {

		return Category{context: list{}}
	}
	var lst list = slice
	return Category{context: lst}
}

//FromList is a list constructor
func FromList(lst list) Category {
	return Category{context: lst}
}

//Map is a Map implementation for a list
func (val list) Map(f Morphism) Category {
	var result list = []interface{}{}
	for _, v := range val {
		result = append(result, f(v))
	}
	return FromList(result)
}

//Map is an Applicative implementation for a list
func (val list) Applicative(jm Category) Category {
	var appl = jm.Value()
	if (appl == empty{}) {
		return Nothing()
	}
	var result list = []interface{}{}
	for _, v := range val {
		result = append(result, appl.(Morphism)(v))
	}
	return FromList(result)
}

//Bind - fmap f Nothing = Nothing
func (val list) Bind(f func(interface{}) Category) Category {
	return f(val)
}

//Value extract a value
func (val list) Value() interface{} {
	return val
}
