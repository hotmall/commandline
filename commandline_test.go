package commandline

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCmdline(t *testing.T) {
	Convey("Test commandline", t, func() {
		Convey("go test -v", func() {
			port := Port()
			So(*h, ShouldBeFalse)
			So(*v, ShouldBeFalse)
			So(isGoTest, ShouldBeTrue)
			So(port, ShouldEqual, 32018)
		})
	})
}
