package functors;

import (
)
import (
	. "github.com/sqdron/go-category"
	"fmt"
)

type Maybe struct {
	val *interface{}
}

var None = Maybe{nil}


func Some(a interface{}) Monad {
	return Maybe{&a}
}

func Nothing() Monad {
	return None;
}

func (m Maybe) Return(value interface{}) Monad {
	return Some(value)
}

func (m Maybe) GetValue() interface{} {
	if m == None {
		return nil
	}
	return *m.val;
}

func (m Maybe) Bind(f func(interface{}, Monad) Monad) Monad {
	if m == None {
		return None
	}
	return f(*m.val, m)
}

func (m Maybe) String() string {
	if m == None {
		return "None"
	}
	return fmt.Sprintf("Some(%v)", *m.val)
}
