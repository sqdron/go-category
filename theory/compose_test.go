package theory;

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Morhism_Composition(t *testing.T) {
	var m1 Morphism = func(ctx interface{}) interface{} {
		return ctx.(int) + 1;
	}

	var m2 Morphism = func(ctx interface{}) interface{} {
		return ctx.(int) + 1;
	};

	var m3 Morphism = func(ctx interface{}) interface{} {
		return ctx.(int) + 1;
	};

	Convey("Make composition", t, func() {
		mResult := m1.Compose(m2, m3);
		r := mResult(1);
		So(r, ShouldNotBeNil);
		So(r, ShouldEqual, 4);
	})
}

func Benchmark_Morphism_Compose(b *testing.B) {
	var m1 Morphism = func(ctx interface{}) interface{} {
		return ctx.(int) + 1;
	}

	var mN Morphism = func(ctx interface{}) interface{} {
		return ctx.(int) + 2;
	};

	for i := 0; i < b.N; i++ {
		r := m1.Compose(mN, mN, mN, mN, mN, mN, mN, mN, mN, mN)
		r(1);
	}
}