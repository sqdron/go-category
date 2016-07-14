package functors

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/smartystreets/assertions/should"
)

func Test_Maybe_Call(t *testing.T) {
	Convey("Just call Maybe func", t, func() {
		result := Maybe(5);
		j := result.(Just);
		So(j.value, should.Equal, 5);
	});
}

func Test_Maybe_Squence_Double(t *testing.T) {
	Convey("Call maybe an apply morphism", t, func() {
		doubleMe := func(i interface{}) interface{} {
			return 2 * i.(int)
		};

		result := Maybe(5).Fmap(doubleMe).Fmap(doubleMe);
		j := result.(Just);
		So(j.value, should.Equal, 20);
	});
}

func Test_Maybe_Squence_Double_And_Nothing(t *testing.T) {
	Convey("Call maybe an apply morphism and one morphism wich return Nil", t, func() {
		doubleMe := func(i interface{}) interface{} {
			return 2 * i.(int)
		};

		nothingBack := func(i interface{}) interface{} {
			return nil;
		};

		result := Maybe(5).Fmap(nothingBack).Fmap(doubleMe);
		So(result, should.Resemble, Nothing{});
	});
}
