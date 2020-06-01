// +build !linux,!freebsd

package udpnofrag

import (
	"net"
)

func UDPSetNoFragment(conn *net.UDPConn) (err error) {
	return
}
