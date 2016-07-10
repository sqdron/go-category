package functors

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type MaybeApplicatives interface {
	Maybe
	Double() MaybeApplicatives;
}

func (j Just) Double() MaybeApplicatives {
	return j.Bind(func (v interface{}) Maybe {
		return Just { 2 * v.(int) }
	}).(MaybeApplicatives);
}

func (n Nothing) Double() MaybeApplicatives {
	return n;
}
func MaybeDouble(a Maybe) Maybe {
	return a.Bind(func (v interface{}) Maybe {
		return Just { 2 * v.(int) }
	})
}

func Test_Maybe_FunctorLaws_Identity(t *testing.T) {
	Convey("Double value several times", t, func() {
		value := Just{}.Return(3);
		result := MaybeDouble(value);
		So(result.GetValue(), ShouldEqual, 6);
		v2 := Just{2}.Double().Double().Double().GetValue();
		So(v2, ShouldEqual, 16);
	})
}

func Test_Maybe_Nothing(t *testing.T) {
	Convey("Double value several times", t, func() {
		value := Nothing{};
		result := MaybeDouble(value);
		So(result.GetValue(), ShouldEqual, Nothing{}.GetValue());
		value2 := Nothing{}.Return(nil);
		result2 := MaybeDouble(value2);
		So(result2.GetValue(), ShouldEqual, Nothing{}.GetValue());
	})
}