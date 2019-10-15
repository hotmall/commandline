package commandline

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestSignal(t *testing.T) {
	suite.Run(t, new(signalSuite))
}

type signalSuite struct {
	suite.Suite
}

func (ss signalSuite) TestSignal() {
	*signal = "kill"
	s, err := Signal()
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), KillSignal, s)
	}

	*signal = "kILl"
	s, err = Signal()
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), KillSignal, s)
	}

	*signal = "KILL"
	s, err = Signal()
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), KillSignal, s)
	}

	*signal = "stop"
	s, err = Signal()
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), StopSignal, s)
	}

	*signal = "StoP"
	s, err = Signal()
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), StopSignal, s)
	}

	*signal = "STOP"
	s, err = Signal()
	assert.Nil(ss.T(), err)
	if assert.NoError(ss.T(), err) {
		assert.Equal(ss.T(), StopSignal, s)
	}

	*signal = "stop1"
	_, err = Signal()
	assert.NotNil(ss.T(), err)
}
