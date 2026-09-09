//go:build linux

package task

import (
        "net"
        "syscall"
)

// bindInterfaceToDialer 为 dialer 设置 SO_BINDTODEVICE 套接字选项
func bindInterfaceToDialer(dialer *net.Dialer, ifaceName string) {
        if ifaceName == "" {
                return
        }
        dialer.Control = func(network, address string, c syscall.RawConn) error {
                var opErr error
                err := c.Control(func(fd uintptr) {
                        opErr = syscall.BindToDevice(int(fd), ifaceName)
                })
                if err != nil {
                        return err
                }
                return opErr
        }
}
