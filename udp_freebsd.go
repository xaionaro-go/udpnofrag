// +build freebsd

package udpnofrag

import (
	"net"
	"syscall"
)

func UDPSetNoFragment(conn *net.UDPConn) (err error) {
	var syscallConn syscall.RawConn
	syscallConn, err = conn.SyscallConn()
	if err != nil {
		return
	}
	err2 := syscallConn.Control(func(fd uintptr) {
		err = syscall.SetsockoptByte(int(fd), syscall.IPPROTO_IP, syscall.IP_DONTFRAG, 1)
	})
	if err != nil {
		return
	}
	err = err2
	return
}
