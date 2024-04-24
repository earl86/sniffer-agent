package session_dealer

import "github.com/earl86/sniffer-agent/model"

type ConnSession interface {
	ReceiveTCPPacket(*model.TCPPacket)
	Close()
}
