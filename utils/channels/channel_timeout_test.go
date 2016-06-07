package channels

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const channelSize = 5

var timeout = time.Duration(500) * time.Millisecond

func TestReceiveTimeout(t *testing.T) {
	ct := NewChannelTimeout(channelSize)
	assert.NotNil(t, ct)

	startTime := time.Now()

	ev, err := ct.Receive(timeout)
	elappsedTime := time.Now().Sub(startTime)

	assert.Nil(t, ev)
	assert.Error(t, err)
	assert.True(t, elappsedTime >= timeout, "Timeout is worng")
}

func TestSendTimeout(t *testing.T) {
	ct := NewChannelTimeout(channelSize)
	assert.NotNil(t, ct)

	for i := 0; i < channelSize; i++ {
		err := ct.Send("event", timeout)
		assert.NoError(t, err)
	}

	startTime := time.Now()
	timeout := time.Duration(500) * time.Millisecond

	err := ct.Send("event", timeout)
	elappsedTime := time.Now().Sub(startTime)

	assert.Error(t, err)
	assert.True(t, elappsedTime >= timeout, "Timeout is worng")
}

func TestReceiveFromClosed(t *testing.T) {
	ct := NewChannelTimeout(channelSize)
	assert.NotNil(t, ct)

	err := ct.Close()
	assert.NoError(t, err)
	ev, err := ct.Receive(0)

	assert.Nil(t, ev)
}

func TestCloseTwice(t *testing.T) {
	ct := NewChannelTimeout(channelSize)
	assert.NotNil(t, ct)

	err := ct.Close()
	assert.NoError(t, err)

	err = ct.Close()
	assert.Error(t, err)
}
