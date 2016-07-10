package functors;

import (
)

type Maybe interface {
	Return(value interface{}) Maybe
	Bind(func(interface{}) Maybe) Maybe
	GetValue() interface{}
}

type Just struct {
	Value interface{}
}

type Nothing struct{}

func (j Just) Return(value interface{}) Maybe {
	return Just{Value:value}
}

func (j Just) GetValue() interface {} {
	return j.Value;
}

func (j Just) Bind(fn func(interface{}) Maybe) Maybe {
	return fn(j.Value)
}

func (n Nothing) Return(value interface{}) Maybe {
	return Nothing{}
}

func (n Nothing) Bind(fn func(interface{}) Maybe) Maybe {
	return Nothing{}
}

func (n Nothing) GetValue() interface {} {
	return nil;
}






