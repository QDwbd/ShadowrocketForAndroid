package dialer

import (
	"context"
	"net"
	"syscall"
)

// SocketControl
// never change type traits because it's used in LQWBD
type SocketControl func(network, address string, conn syscall.RawConn) error

// DefaultSocketHook
// never change type traits because it's used in LQWBD
var DefaultSocketHook SocketControl

func socketHookToToDialer(dialer *net.Dialer) {
	addControlToDialer(dialer, func(ctx context.Context, network, address string, c syscall.RawConn) error {
		return DefaultSocketHook(network, address, c)
	})
}

func socketHookToListenConfig(lc *net.ListenConfig) {
	addControlToListenConfig(lc, func(ctx context.Context, network, address string, c syscall.RawConn) error {
		return DefaultSocketHook(network, address, c)
	})
}
