package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

var (
	ssid     string
	password string
)

var (
	securityType = uint32(0x00400004)
)

func main() {
	err := run()
	for err != nil {
		fmt.Printf("error: %s\r\n", err.Error())
		time.Sleep(5 * time.Second)
	}
}

var buf [1024]byte

func run() error {
	rtl, err := setupRTL8720DN()
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_wifi_off()
	if err != nil {
		return err
	}
	_, err = rtl.Rpc_wifi_on(0x00000001)
	if err != nil {
		return err
	}

	// WiFi.disconnect();
	_, err = rtl.Rpc_wifi_disconnect()
	if err != nil {
		return err
	}

	// delay(100);
	time.Sleep(100 * time.Millisecond)

	for i := 0; i < 5; i++ {
		fmt.Printf("Connecting to WiFi..\r\n")
		ret, err := rtl.Rpc_wifi_connect(ssid, password, securityType, -1, 0)
		if err != nil {
			return err
		}
		if ret != 0 {
			fmt.Printf("wifi connect failed\r\n")
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	_, err = rtl.Rpc_tcpip_adapter_dhcpc_start(0)
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		_, err = rtl.Rpc_wifi_is_connected_to_ap()
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	ip_info := make([]byte, 12)
	_, err = rtl.Rpc_tcpip_adapter_get_ip_info(0, &ip_info)
	if err != nil {
		return err
	}
	ip := newIPInfo(ip_info)
	ip.Dump()

	for {
		time.Sleep(1 * time.Hour)
	}

	return nil
}

type ipInfo struct {
	IP      uint32
	Mask    uint32
	Gateway uint32
}

func newIPInfo(b []byte) *ipInfo {
	return &ipInfo{
		IP:      binary.BigEndian.Uint32(b),
		Mask:    binary.BigEndian.Uint32(b[4:]),
		Gateway: binary.BigEndian.Uint32(b[8:]),
	}
}

func (i *ipInfo) Dump() {
	fmt.Printf("IP Address : %d.%d.%d.%d\r\n", byte(i.IP>>24), byte(i.IP>>16), byte(i.IP>>8), byte(i.IP))
	fmt.Printf("Mask       : %d.%d.%d.%d\r\n", byte(i.Mask>>24), byte(i.Mask>>16), byte(i.Mask>>8), byte(i.Mask))
	fmt.Printf("Gateway    : %d.%d.%d.%d\r\n", byte(i.Gateway>>24), byte(i.Gateway>>16), byte(i.Gateway>>8), byte(i.Gateway))
}
