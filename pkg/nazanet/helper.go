// Copyright 2020, Chef.  All rights reserved.
// https://github.com/souliot/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package nazanet

import (
	"fmt"
	"net"
)

func listenUDPWithPort(port uint16) (*net.UDPConn, error) {
	addr := fmt.Sprintf(":%d", port)
	return listenUDPWithAddr(addr)
}

func listenUDPWithAddr(addr string) (*net.UDPConn, error) {
	udpAddr, err := net.ResolveUDPAddr(udpNetwork, addr)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP(udpNetwork, udpAddr)
}
