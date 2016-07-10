package functors

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	. "github.com/sqdron/go-category"
	"github.com/smartystreets/assertions/should"
)

//type IntMonad int;
//
//func (m IntMonad) Bind(f func(interface{}, Monad) Monad) Monad {
//	result := IntMonad{}
//	for _, val := range m {
//		next := f(val, m).(IntMonad)
//		result = append(result, next...)
//	}
//	return result
//}
//
//func (m IntMonad) Return(a interface{}) Monad {
//	return IntMonad{a.(int)}
//}


//
//type MathMonad interface {
//	Monad
//	Add(value int) MathMonad;
//	Double() MathMonad;
//}
//
//func (j Just) ToMath() MathMonad {
//	return Just{}.Return(j.Value).(MathMonad);
//}
//
//func (j Nothing) ToMath() MathMonad {
//	return Nothing{}.Return(nil).(MathMonad);
//}
//
//func DoubleValue(m Monad) Monad {
//	return m.Bind(func(v interface{}) Monad {
//		return Just{2 * v.(int) }
//	});
//}
//
//func AddValue(m Monad, value int) Monad {
//	return m.Bind(func(v interface{}) Monad {
//		return Just{ value + v.(int) }
//	});
//}
//
//func (j Just) Double() MathMonad {
//	return DoubleValue(j).(MathMonad);
//}
//
//func (j Just) Add(value int) MathMonad {
//	return AddValue(j, value).(MathMonad);
//}
//
//
//func (n Nothing) Double() MathMonad {
//	return DoubleValue(n).(MathMonad);
//}
//
//func (n Nothing) Add(value int) MathMonad {
//	return AddValue(n, value).(MathMonad);
//}
//
//

type IntSliceMonad []int

func (m IntSliceMonad) Bind(f func(interface{}, Monad) Monad) Monad {
	result := IntSliceMonad{}
	for _, val := range m {
		next := f(val, m).(IntSliceMonad)
		result = append(result, next...)
	}
	return result
}

func (m IntSliceMonad) Return(a interface{}) Monad {
	return IntSliceMonad{a.(int)}
}

func (m IntSliceMonad) GetValue() interface{} {
	return m;
}

func Test_Maybe_Functor_Double(t *testing.T) {
	Convey("Double value several times", t, func() {
		doubleMe := func(i interface{}, m Monad) Monad {
			return m.Return(2 * i.(int))
		}
		value := Some(3).Bind(doubleMe).Bind(doubleMe);
		So(value.GetValue(), ShouldEqual, 12);
	})
}

func Test_Maybe_Functor_Slices_Double(t *testing.T) {
	Convey("Double value int slice times", t, func() {
		doubleMe := func(i interface{}, m Monad) Monad {
			return m.Return(2 * i.(int))
		}

		var oneTwoThree = IntSliceMonad{1, 2, 3}
		r := oneTwoThree.Bind(doubleMe).GetValue().(IntSliceMonad);
		So(r[0], ShouldEqual, 2);
		So(r[1], ShouldEqual, 4);
		So(r[2], ShouldEqual, 6);
		//So(r([]int)[0],ShouldEqual, 2);
	})
}

func Test_Maybe_Functor_Double_And_None(t *testing.T) {
	Convey("Double value several times", t, func() {
		doubleMe := func(i interface{}, m Monad) Monad {
			return m.Return(2 * i.(int))
		}

		oops := func(i interface{}, m Monad) Monad {
			return None
		}

		value := Some(3).Bind(doubleMe).Bind(oops).Bind(doubleMe);
		So(value.GetValue(), ShouldEqual, None.GetValue());
	})
}

//func Test_Maybe_Functor_Slices_Double_And_None(t *testing.T) {
//	Convey("Double value int slice times", t, func() {
//		doubleMe := func(i interface{}, m Monad) Monad {
//			return m.Return(2 * i.(int))
//		}
//
//		oops := func(i interface{}, m Monad) Monad {
//			return None
//		}
//
//		var oneTwoThree = IntSliceMonad{1, 2, 3}
//		r := oneTwoThree.Bind(doubleMe).Bind(oops).Bind(doubleMe).GetValue().(IntSliceMonad);
//		So(r[0], ShouldEqual, 4);
//		So(r[1], ShouldEqual, 8);
//		So(r[2], ShouldEqual, 12);
//		//So(r([]int)[0],ShouldEqual, 2);
//	})
//}

//func Test_Maybe_Functor_Double_And_Add(t *testing.T) {
//	Convey("Double value several times", t, func() {
//		value := Just{2}.ToMath().Double().Add(10);
//		So(value.GetValue(), ShouldEqual, 14);
//	})
//}

func Test_Maybe_Nothing(t *testing.T) {
	Convey("Double Nothing", t, func() {

		doubleMe := func(i interface{}, m Monad) Monad {
			return m.Return(2 * i.(int))
		}
		value := Nothing().Bind(doubleMe);
		So(value.GetValue(), should.Equal, None.GetValue());
	})
}