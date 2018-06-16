package shhext

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/status-im/status-go/signal"
)

// EnvelopeSignalHandler sends signals when envelope is sent or expired.
type EnvelopeSignalHandler struct{}

// EnvelopeSent triggered when envelope delivered atleast to 1 peer.
func (h EnvelopeSignalHandler) EnvelopeSent(hash common.Hash) {
	signal.SendEnvelopeSent(hash)
}

// EnvelopeExpired triggered when envelope is expired but wasn't delivered to any peer.
func (h EnvelopeSignalHandler) EnvelopeExpired(hash common.Hash) {
	signal.SendEnvelopeExpired(hash)
}

// MailServerRequestCompleted triggered when mailserver send a ack with a requesID sent previously
func (h EnvelopeSignalHandler) MailServerRequestCompleted(hash common.Hash) {
	signal.SendMailServerRequestCompleted(hash)
}

// MailServerRequestExpired triggered when mail server request expires
func (h EnvelopeSignalHandler) MailServerRequestExpired(hash common.Hash) {
	signal.SendMailServerRequestExpired(hash)
}
