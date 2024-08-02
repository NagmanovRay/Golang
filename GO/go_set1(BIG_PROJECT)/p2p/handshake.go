package p2p

import "errors"

// Error msg
var ErrInvaklidHandShake = errors.New("invalid error")

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
