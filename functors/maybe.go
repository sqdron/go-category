package functors;

import (
	. "github.com/sqdron/go-category"
)

type Just struct {
	value interface{}
}

type Nothing struct {}

type IMaybe interface {
	Fmap(m Morphism) IMaybe;
}

func Maybe(val interface{}) IMaybe{
	if (val == nil){
		return Nothing{};
	}
	return Just{value:val}
}

func (n Nothing) Fmap(m Morphism)IMaybe{
	return Nothing{};
}

func (n Just) Fmap(fn Morphism) IMaybe{
	return Maybe(fn(n.value));
}
