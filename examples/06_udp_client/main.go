package main

import (
	"encoding/binary"
	"fmt"
	"time"
)

var (
	ssid     string
	password string
	server   string
	port     int
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

	_, err = rtl.Rpc_wifi_disconnect()
	if err != nil {
		return err
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Connecting to WiFi network\r\n")

	for i := 0; i < 5; i++ {
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
	fmt.Printf("WiFi connected!\r\n")

	ip_info := make([]byte, 12)
	_, err = rtl.Rpc_tcpip_adapter_get_ip_info(0, &ip_info)
	if err != nil {
		return err
	}
	ip := newIPInfo(ip_info)
	ip.Dump()

	socket, err := rtl.Rpc_lwip_socket(0x02, 0x02, 0x00)
	if err != nil {
		return err
	}

	optval := []byte{0x01, 0x00, 0x00, 0x00}
	_, err = rtl.Rpc_lwip_setsockopt(socket, 0x00000FFF, 0x00000004, optval, uint32(len(optval)))
	if err != nil {
		return err
	}

	// address
	name := []byte{0x00, 0x02, 0x0D, 0x05, 0xC0, 0xA8, 0x01, 0x78, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	name[2] = byte(port >> 8)
	name[3] = byte(port)
	name[4] = byte(ip.IP >> 24)
	name[5] = byte(ip.IP >> 16)
	name[6] = byte(ip.IP >> 8)
	name[7] = byte(ip.IP)
	_, err = rtl.Rpc_lwip_bind(socket, name, uint32(len(name)))
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_fcntl(socket, 0x00000004, 0x00000000)
	if err != nil {
		return err
	}

	seconds := 0
	for {
		addr := make([]byte, 4)
		_, err = rtl.Rpc_netconn_gethostbyname(server, &addr)
		if err != nil {
			return err
		}
		fmt.Printf("%#v\r\n", addr)

		dataptr := []byte(fmt.Sprintf("Seconds since bool: %d", seconds))
		to := []byte{0x00, 0x02, 0x0D, 0x05, 0xC0, 0xA8, 0x01, 0x76, 0x2F, 0x44, 0x00, 0x00, 0x22, 0x2A, 0x01, 0x00}
		rtl.Rpc_lwip_sendto(socket, dataptr, 0x00000000, to, uint32(len(to)))
		time.Sleep(1 * time.Second)
		seconds++
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
