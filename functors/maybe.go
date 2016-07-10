package functors;

import (
)
import . "github.com/sqdron/go-category"

type Just struct {
	Value interface{}
}

type Nothing struct{}

func (j Just) Return(value interface{}) Monad {
	return Just{Value:value}
}

func (j Just) GetValue() interface {} {
	return j.Value;
}

func (j Just) Bind(fn func(interface{}) Monad) Monad {
	return fn(j.Value)
}

func (n Nothing) Return(value interface{}) Monad {
	return Nothing{}
}

func (n Nothing) Bind(fn func(interface{}) Monad) Monad {
	return Nothing{}
}

func (n Nothing) GetValue() interface {} {
	return nil;
}






