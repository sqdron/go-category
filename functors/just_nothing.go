package functors

type Just struct {
	value interface{}
}

type Nothing struct {}

func J(value interface{}) Just {
	return Just{value:value};
}

func N() Nothing {
	return Nothing{};
}