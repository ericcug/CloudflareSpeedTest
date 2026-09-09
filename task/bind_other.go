//go:build !linux

package task

import "net"

func bindInterfaceToDialer(dialer *net.Dialer, ifaceName string) {
        // 非 Linux 系统不支持 SO_BINDTODEVICE，保持原有 LocalAddr 行为或留空
}
