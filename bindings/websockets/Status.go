//go:build wasm

package websockets

type Status int

const (

	// https://www.rfc-editor.org/rfc/rfc6455.html

	StatusNormalClosure           Status = 1000
	StatusGoingAway               Status = 1001
	StatusProtocolError           Status = 1002
	StatusUnsupportedData         Status = 1003
	status_undefined_1004         Status = 1004 // Reserved; must not be sent.
	StatusNoStatusReceived        Status = 1005 // Reserved; must not be sent.
	StatusAbnormalClosure         Status = 1006 // Reserved; must not be sent.
	StatusInvalidFramePayloadData Status = 1007
	StatusPolicyViolation         Status = 1008
	StatusMessageTooBig           Status = 1009
	StatusMandatoryExtension      Status = 1010 // Client only.
	StatusInternalServerError     Status = 1011

	StatusServiceRestart    Status = 1012
	StatusTryAgainLater     Status = 1013
	StatusBadGateway        Status = 1014
	StatusTLSHandshakeError Status = 1015 // Reserved; must not be sent.

)

