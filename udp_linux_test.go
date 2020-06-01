// +build linux testrareexceptions

package udpnofrag

import (
	"net"
	"reflect"
	"syscall"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/xaionaro-go/unsafetools"
)

func TestUdpSetNoFragment_negative(t *testing.T) {
	assert.Equal(t, syscall.EINVAL, UDPSetNoFragment(&net.UDPConn{}))
	conn, err := net.Dial("udp", "127.0.0.1:1")
	udpConn := conn.(*net.UDPConn)
	assert.NoError(t, err)
	fd := (**struct{ a [65536]byte })((unsafe.Pointer)(reflect.ValueOf(unsafetools.FieldByName(udpConn, `fd`)).Elem().UnsafeAddr()))
	*fd = &struct{ a [65536]byte }{}
	assert.Equal(t, syscall.ENOTSOCK, UDPSetNoFragment(udpConn))
}
