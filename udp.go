package udpnofrag

import (
	"net"
)

// UDPSetNoFragment sets flag on the socket of UDP connection `conn` to do
// not fragment payloads.
func UDPSetNoFragment(conn *net.UDPConn) error {
	return udpSetNoFragment(conn)
}
