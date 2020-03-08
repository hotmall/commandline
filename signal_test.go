package commandline

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSignal(t *testing.T) {
	Convey("Test signal", t, func() {
		Convey("kill", func() {
			*signal = "kill"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, KillSignal)
		})

		Convey("kILl", func() {
			*signal = "kILl"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, KillSignal)
		})

		Convey("KILL", func() {
			*signal = "KILL"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, KillSignal)
		})

		Convey("stop", func() {
			*signal = "stop"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, StopSignal)
		})

		Convey("StoP", func() {
			*signal = "StoP"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, StopSignal)
		})

		Convey("STOP", func() {
			*signal = "STOP"
			s, err := Signal()
			So(err, ShouldBeNil)
			So(s, ShouldEqual, StopSignal)
		})

		Convey("stop1", func() {
			*signal = "stop1"
			s, err := Signal()
			So(err.Error(), ShouldEqual, `Unrecognized signal: "stop1"`)
			// Unrecognized signal repreasents StopSignal
			So(s, ShouldEqual, StopSignal)
		})
	})
}
