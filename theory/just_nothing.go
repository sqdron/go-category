package theory;

type Just struct {
	value interface{}
}

type Nothing struct {}

type Morphism func(interface{}) interface{};

func J(value interface{}) Just {
	return Just{value:value};
}

func N() Nothing {
	return Nothing{};
}