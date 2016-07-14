package functors

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
	"github.com/sqdron/go-category"
)

func Test_Maybe_Applicative_Sum(t *testing.T) {
	Convey("Just call Maybe with sum applicative", t, func() {
		var add category.Morphism = func(value interface{}) interface{} {
			var result category.Morphism = func(value2 interface{}) interface{} {
				return value.(int) + value2.(int);
			}
			return result;
		}

		result := Maybe(add).A(J(4)).A(J(2));
		So(result.(Just).value, should.Equal, 6);
	});
}

func Test_Maybe_Applicative_Add_3(t *testing.T) {
	Convey("Just call Maybe with +3 applicative", t, func() {
		var add category.Morphism = func(value interface{}) interface{} {
			return value.(int) + 3;
		}

		result := Maybe(add).A(J(4));
		So(result.(Just).value, should.Equal, 7);
	});
}

func Test_Maybe_Applicative_Multiplicate(t *testing.T) {
	Convey("Just call Maybe with multiplication applicative", t, func() {
		var mult category.Morphism = func(value interface{}) interface{} {
			var result category.Morphism = func(value2 interface{}) interface{} {
				return value.(int) * value2.(int);
			}
			return result;
		}

		result := Maybe(mult).A(J(4)).A(J(3));
		So(result.(Just).value, should.Equal, 12);
	});
}


func Test_Maybe_Applicative_Multiplicate_With_Nothing(t *testing.T) {
	Convey("Just call Maybe with multiplication applicative and nothing", t, func() {
		var mult category.Morphism = func(value interface{}) interface{} {
			var result category.Morphism = func(value2 interface{}) interface{} {
				return value.(int) * value2.(int);
			}
			return result;
		}

		result := Maybe(mult).A(J(4)).A(N());
		So(result, should.Resemble, Nothing{});
	});
}

