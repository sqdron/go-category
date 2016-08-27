package helpers

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Slices(t *testing.T) {
	var m1 interface{} = []int{1, 2, 3}

	Convey("Make slice of interfaces", t, func() {
		mResult := InterfaceSlice(m1)
		So(len(mResult), ShouldEqual, 3)
	})
}
