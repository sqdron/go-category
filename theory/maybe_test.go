package theory;

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
)

func Test_Functor_Maybe(t *testing.T) {
	var addFiveFunc Morphism = func(i interface{}) interface{} {
		return 5 + i.(int)
	};

	var doubleFunc Morphism = func(i interface{}) interface{} {
		return 2 * i.(int)
	};

	Convey("Construct Maybe", t, func() {
		j5 := Just(5);
		So(j5.Value(), should.Equal, 5);
		Convey("And Double it", func() {
			j10 := j5.FMap(doubleFunc);
			So(j10.Value(), should.Equal, 10);

			Convey("And now Add 5", func() {
				j15 := j10.FMap(addFiveFunc);
				So(j15.Value(), should.Equal, 15);
				Convey("And now Fmap in haskel Style", func() {
					result := addFiveFunc.S()(j15);
					So(result.Value(), should.Equal, 20);
					Convey("Now time for Nothing", func() {
						n := addFiveFunc.S().Nothing();
						So(n.Value(), should.Resemble, empty{});
					})
				});
			});
		});
	});
}

func Test_Applicative_Maybe(t *testing.T) {
	var doubleFunc Morphism = func(i interface{}) interface{} {
		return 2 * i.(int)
	};

	var addFunc Morphism = func(value interface{}) interface{} {
		var result Morphism = func(value2 interface{}) interface{} {
			return value.(int) + value2.(int);
		}
		return result;
	}

	Convey("Applicative test double 5", t, func() {
		j10 := Just(doubleFunc).A().Just(5);
		So(j10.Value(), should.Equal, 10);

		Convey("Add two values", func() {
			n := addFunc.S().Just(2).A().Just(2)
			So(n.Value(), should.Resemble, 4);
			Convey("And Double nothing", func() {
				n := Just(doubleFunc).A().Nothing();
				So(n.Value(), should.Resemble, empty{});
				Convey("And Two values in Haskel style", func() {
					n := addFunc.S().Just(8).A().Just(2);
					So(n.Value(), should.Equal, 10);
				});
				Convey("And Add With Nothing", func() {
					addFunc.S().Nothing().A().Just(2);
					So(n.Value(), should.Resemble, empty{});
				});
				Convey("And Add With Nothing Right", func() {
					addFunc.S().Just(2).A().Nothing();
					So(n.Value(), should.Resemble, empty{});
				});
			});
		});


	});
}

func Test_Monad_Maybe(t *testing.T) {
	addMonadf := func(i interface{}) Category {
		return Maybe(i.(int) + 5);
	};

	Convey("Construct Maybe", t, func() {
		j5 := Just(5);
		So(j5.Value(), should.Equal, 5);
		Convey("Add by monad", func() {
			j10 := j5.Bind(addMonadf);
			So(j10.Value(), should.Equal, 10);
		});

		Convey("Add by monad with nothing", func() {
			So(Maybe(nil).Bind(addMonadf).Value(), should.Resemble, empty{});
		});
	});
}