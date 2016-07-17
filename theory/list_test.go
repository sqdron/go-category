package theory

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
	"github.com/sqdron/go-category/helpers"
)

func Test_Monad_List(t *testing.T) {
	var doubleFunc Morphism = func(i interface{}) interface{} {
		return 2 * i.(int)
	};

	var addFiveFunc Morphism = func(i interface{}) interface{} {
		return 5 + i.(int)
	};

	Convey("Construct Maybe", t, func() {
		j5 := List([]int{1, 3, 5, 7});
		So(isEqual(j5.Value(), []int{1, 3, 5, 7}), should.BeTrue);
		Convey("And Double it", func() {
			j10 := j5.FMap(doubleFunc);
			So(isEqual(j10.Value(), []int{2, 6, 10, 14}), should.BeTrue);
			Convey("And now Add 5", func() {
				j15 := j10.FMap(addFiveFunc);
				So(isEqual(j15.Value(), []int{7, 11, 15, 19}), should.BeTrue);
				Convey("And now Fmap in haskel Style", func() {
					result := addFiveFunc.S()(j15);
					So(isEqual(result.Value(), []int{12, 16, 20, 24}), should.BeTrue);
					Convey("Now time for Nothing", func() {
						n := addFiveFunc.S().Nothing();
						So(n.Value(), should.Resemble, empty{});
					})
				});
			});
		});
	});
}

//TODO: make a library function for comparing lists and slices
func isEqual(slice1 interface{}, slice2 interface{}) bool {
	s1 := helpers.InterfaceSlice(slice1);
	s2 := helpers.InterfaceSlice(slice2);
	if (len(s1) != len(s2)) {
		return false;
	}

	for i := range s1 {
		if (s1[i] != s2[i]) {
			return false;
		}
	}
	return true
}