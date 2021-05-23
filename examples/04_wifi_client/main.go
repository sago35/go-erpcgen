package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ssid     string
	password string
	hostname string
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
	_, err = rtl.Rpc_tcpip_adapter_get_ip_info(0, ip_info)
	if err != nil {
		return err
	}
	ip := newIPInfo(ip_info)
	ip.Dump()

	addr := make([]byte, 15)
	_, err = rtl.Rpc_dns_gethostbyname_addrtype(hostname, addr, 0x00004d4d, []byte{0x00, 0x08, 0x00, 0x00, 0x00, 0x4d, 0x4d, 0x00, 0x00, 0xB0, 0xD6, 0x00, 0x20}, 0x02)
	if err != nil {
		return err
	}
	//fmt.Printf("%#v\r\n", addr)
	ipaddr := ParseIP(string(addr))
	port := 80

	// TODO: rpc_wifi_dns_found()

	socket, err := rtl.Rpc_lwip_socket(0x02, 0x01, 0x00)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_fcntl(socket, 0x00000003, 0x00000000)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_fcntl(socket, 0x00000004, 0x00000001)
	if err != nil {
		return err
	}

	name := []byte{0x00, 0x02, 0x00, 0x50, 0xC0, 0xA8, 0x01, 0x76, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if true {
		name[3] = byte(port)
		name[4] = byte(ipaddr[0])
		name[5] = byte(ipaddr[1])
		name[6] = byte(ipaddr[2])
		name[7] = byte(ipaddr[3])
	}
	_, err = rtl.Rpc_lwip_connect(socket, name, uint32(len(name)))
	if err != nil {
		return err
	}

	readset := []byte{}
	writeset := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	exceptset := []byte{}
	timeout := []byte{}
	_, err = rtl.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	if err != nil {
		return err
	}

	optlen := uint32(4)
	_, err = rtl.Rpc_lwip_getsockopt(socket, 0x00000FFF, 0x00001007, []byte{0xA5, 0xA5, 0xA5, 0xA5}, nil, optlen)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_fcntl(socket, 0x00000003, 0x00000000)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_fcntl(socket, 0x00000004, 0x00000000)
	if err != nil {
		return err
	}

	readset = []byte{}
	writeset = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	exceptset = []byte{}
	timeout = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x42, 0x0F, 0x00, 0xFF, 0xFF, 0xFF, 0xFF}
	_, err = rtl.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.1\nHost: "+hostname+"\n\n"), 0x00000008)
	if err != nil {
		return err
	}

	//rtl.Debug(true)
	time.Sleep(1 * time.Second)
	rcvBuf := make([]byte, 0x0400)
	n, err := rtl.Rpc_lwip_recv(socket, rcvBuf, uint32(len(rcvBuf)), 0x00000008, 0x00002800)
	if err != nil {
		return err
	}
	header := httpHeader(rcvBuf[:n])
	length := header.ContentLength()
	if n > 0 {
		fmt.Print(string(rcvBuf[:n]))
	}

	for length > 0 {
		n, err := rtl.Rpc_lwip_recv(socket, rcvBuf, uint32(len(rcvBuf)), 0x00000008, 0x00002800)
		if err != nil {
			return err
		}
		length -= int(n)
		if n > 0 {
			fmt.Print(string(rcvBuf[:n]))
		} else if n < 0 {
			length = 0
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}

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

type httpHeader []byte

func (h httpHeader) ContentLength() int {
	contentLength := -1
	idx := bytes.Index(h, []byte("Content-Length: "))
	if 0 <= idx {
		_, err := fmt.Sscanf(string(h[idx+16:]), "%d", &contentLength)
		if err != nil {
			return -1
		}
	}
	return contentLength
}

type IP []byte

// ParseIP parses s as an IP address, returning the result.
func ParseIP(s string) IP {
	ret := make([]byte, 4)

	sp := strings.Split(strings.TrimRight(s, "\x00"), ".")
	fmt.Printf("%#v\r\n", sp)
	x, _ := strconv.ParseUint(sp[0], 10, 8)
	ret[0] = byte(x)

	x, _ = strconv.ParseUint(sp[1], 10, 8)
	ret[1] = byte(x)

	x, _ = strconv.ParseUint(sp[2], 10, 8)
	ret[2] = byte(x)

	x, _ = strconv.ParseUint(sp[3], 10, 8)
	ret[3] = byte(x)

	return IP([]byte(ret))
}

// String returns the string form of the IP address ip.
func (ip IP) String() string {
	return string(ip)
}
