package commandline

import (
	"testing"

	. "gopkg.in/check.v1"
)

func TestSignal(t *testing.T) { TestingT(t) }

type signalSuite struct{}

var _ = Suite(&signalSuite{})

func (ss signalSuite) TestSignal(c *C) {
	_, err := Signal()
	c.Check(err, NotNil)
	c.Check(err, ErrorMatches, `Unrecognized signal: ""`)

	*signal = "stop"
	s, err := Signal()
	c.Check(err, IsNil)
	c.Check(s, Equals, StopSignal)

	*signal = "STOP"
	s, err = Signal()
	c.Check(err, IsNil)
	c.Check(s, Equals, StopSignal)

	*signal = "stop1"
	_, err = Signal()
	c.Check(err, ErrorMatches, `Unrecognized signal: "stop1"`)
}
