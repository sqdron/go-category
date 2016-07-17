package theory;

import "github.com/sqdron/go-category/helpers"

type list []interface{};

type Ilist interface {
	IMonad
}

//TODO: replace on arguments range ...items
func List(slice interface{}) Category {
	var lst list = list{};
	if (slice != nil){
		lst = helpers.InterfaceSlice(slice);
	};
	return Category{context:lst};
}

func FromList(lst list) Category {
	return Category{context:lst};
}

func (val list) Map(f Morphism) Category{
	//TODO: allocate memory
	var result list = []interface{}{};
	for _, v := range val {
		result = append(result, f(v))
	}
	//return FromList(result);
	return FromList(result);
}

func (val list) Applicative(jm Category) Category{
	var appl = jm.Value();
	if (appl == empty{}) {
		return Nothing()
	}
	var f Morphism = appl.(Morphism);
	//TODO: allocate memory
	var result list = []interface{}{};
	for _, v := range val {
		result = append(result, f(v))
	}
	return FromList(result);
}

// fmap f Nothing = Nothing
func (val list) Bind(f func(interface{}) Category) Category{
	return f(val);
}

func (val list) Value() interface {}{
	return val;
}
