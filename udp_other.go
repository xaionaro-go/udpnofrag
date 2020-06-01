// +build !linux,!freebsd

package udpnofrag

import (
	"net"
)

func udpSetNoFragment(conn *net.UDPConn) (err error) {
	return
}
