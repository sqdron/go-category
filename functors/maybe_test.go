package functors

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/sqdron/go-category"
)

type MathMonad interface {
	Monad
	Add(value int) MathMonad;
	Double() MathMonad;
}

func (j Just) ToMath() MathMonad {
	return Just{}.Return(j.Value).(MathMonad);
}

func (j Nothing) ToMath() MathMonad {
	return Nothing{}.Return(nil).(MathMonad);
}

func DoubleValue(m Monad) Monad {
	return m.Bind(func(v interface{}) Monad {
		return Just{2 * v.(int) }
	});
}

func AddValue(m Monad, value int) Monad {
	return m.Bind(func(v interface{}) Monad {
		return Just{ value + v.(int) }
	});
}

func (j Just) Double() MathMonad {
	return DoubleValue(j).(MathMonad);
}

func (j Just) Add(value int) MathMonad {
	return AddValue(j, value).(MathMonad);
}


func (n Nothing) Double() MathMonad {
	return DoubleValue(n).(MathMonad);
}

func (n Nothing) Add(value int) MathMonad {
	return AddValue(n, value).(MathMonad);
}


func Test_Maybe_Functor_Double(t *testing.T) {
	Convey("Double value several times", t, func() {
		value := Just{2}.ToMath().Double();
		//result := MaybeDouble(value);
		So(value.GetValue(), ShouldEqual, 4);
		//v2 := Just{2}.Double().Double().Double().GetValue();
		//So(v2, ShouldEqual, 16);
	})
}

func Test_Maybe_Functor_Double_And_Add(t *testing.T) {
	Convey("Double value several times", t, func() {
		value := Just{2}.ToMath().Double().Add(10);
		So(value.GetValue(), ShouldEqual, 14);
	})
}


func Test_Maybe_Nothing(t *testing.T) {
	Convey("Double value several times", t, func() {
		value := Nothing{}.ToMath().Double();
		//result := MaybeDouble(value);
		So(value.GetValue(), ShouldEqual, Nothing{}.GetValue());
	})
}