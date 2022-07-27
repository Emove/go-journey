package channel

import "testing"

func TestNoBuffChanTest(t *testing.T) {
	NoBuffChanTest()
}

func TestSendToNilChannel(t *testing.T) {
	SendToNilChannel()
}

func TestReceiveFromNilChannel(t *testing.T) {
	ReceiveFromNilChannel()
}

func TestSendToClosedChannel(t *testing.T) {
	SendToClosedChannel()
}

func TestCloseAClosedChannel(t *testing.T) {
	CloseAClosedChannel()
}

func TestReceiveFromEmptyChannel(t *testing.T) {
	ReceiveFromEmptyChannel()
}
