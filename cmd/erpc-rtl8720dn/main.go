package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/sago35/go-erpcgen/rtl8720dn"

	"go.bug.st/serial"
)

var (
	seq         uint64
	adapterInit bool
	reconnect   bool
	port        string
	ssid        string
	password    string
)

var (
	securityType = uint32(0x00400004)
)

func main() {
	flag.Uint64Var(&seq, "seq", 1, "seq")
	flag.BoolVar(&adapterInit, "init", false, "init")
	flag.BoolVar(&reconnect, "reconnect", false, "init")
	flag.StringVar(&port, "port", "", "port")
	flag.Parse()

	if adapterInit {
		reconnect = true
		seq = 1
	}

	err := run(port, seq)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("end...\n")
	//for {
	//	time.Sleep(1 * time.Second)
	//}
}

var (
	p serial.Port
)

func run(port string, seq uint64) error {
	// request := createRequest()
	// startWriteMessage(type, service, request, sequence)
	// performRequest(request)
	//   この中で応答の header + sequence の 4byte を読み込む
	// read(&result)
	var err error
	p, err = serial.Open(port, &serial.Mode{BaudRate: 115200})
	if err != nil {
		return err
	}

	rtl := rtl8720dn.New(p)

	rtl.SetSeq(seq)

	go rtl.ReadThread()

	if adapterInit {
		_, err = rtl.Rpc_tcpip_adapter_init()
		if err != nil {
			return err
		}
	}

	if reconnect {
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

		ret, err := rtl.Rpc_wifi_connect(ssid, password, securityType, -1, 0)
		if err != nil {
			return err
		}
		if ret != 0 {
			return fmt.Errorf("wifi connect failed\n")
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
	}

	ip_info := make([]byte, 12)
	_, err = rtl.Rpc_tcpip_adapter_get_ip_info(0, ip_info)
	if err != nil {
		return err
	}
	ip := newIPInfo(ip_info)
	ip.Dump()

	//	addr := make([]byte, 4)
	//	_, err = rtl.Rpc_dns_gethostbyname_addrtype(`192.168.1.110`, addr, 0x00006d61, []byte{0x61, 0x6d, 0x00, 0x00, 0x38, 0x94, 0x00, 0x20}, 0x02)
	//	//_, err = rtl.Rpc_dns_gethostbyname_addrtype(`tinygo.org`, addr, 0x00006d61, []byte{0x61, 0x6d, 0x00, 0x00, 0x38, 0x94, 0x00, 0x20}, 0x02)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("addr : %#v\n", addr)
	//
	//	socket, err := rtl.Rpc_lwip_socket(0x02, 0x01, 0x00)
	//	if err != nil {
	//		return err
	//	}
	//
	//	name := []byte{0x00, 0x02, 0x00, 0x50, 0xc0, 0xa8, 0x01, 0x6e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//	_, err = rtl.Rpc_lwip_connect(socket, name, uint32(len(name)))
	//	if err != nil {
	//		return err
	//	}
	//
	//	readset := []byte{}
	//	writeset := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//	exceptset := []byte{}
	//	timeout := []byte{}
	//	_, err = rtl.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	//	if err != nil {
	//		return err
	//	}
	//
	//	optlen := uint32(4)
	//	_, err = rtl.Rpc_lwip_getsockopt(socket, 0x00000FFF, 0x00001007, []byte{0xA5, 0xA5, 0xA5, 0xA5}, nil, &optlen)
	//	if err != nil {
	//		return err
	//	}
	//
	//	{
	//		readset := []byte{}
	//		writeset := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	//		exceptset := []byte{}
	//		timeout := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x42, 0x0f, 0x00, 0xff, 0xff, 0xff, 0xff}
	//		_, err = rtl.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	//		if err != nil {
	//			return err
	//		}
	//	}
	//
	//	//_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.0\n\n"), 0x00000008)
	//	//_, err = rtl.Rpc_lwip_send(socket, []byte("GET / HTTP/1.1\n\n"), 0x00000008)
	//	_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.1\nHost: 192.168.1.110\n\n"), 0x00000008)
	//	//_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.1\n\n"), 0x00000008)
	//	//_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.1\nHost: 192.168.1.110\nUser-Agent: curl/7.68.0\nAccept: */*\nConnection: close\n\n"), 0x00000008)
	//	//_, err = rtl.Rpc_lwip_send(socket, []byte("GET /index.html HTTP/1.1\nHost: 192.168.1.110\n\n"), 0x00000008)
	//	if err != nil {
	//		return err
	//	}
	//
	//	rcvBuf := make([]byte, 0x0400)
	//	rcvLen, err := rtl.Rpc_lwip_recv(socket, rcvBuf, uint32(len(rcvBuf)), 0x00000008, 0x00002800)
	//	if err != nil {
	//		return err
	//	}
	//
	//	if rcvLen >= 0 {
	//		fmt.Printf("rcv:\n--\n%s\n--\n", string(rcvBuf[:rcvLen]))
	//	}
	//
	//	_, err = rtl.Rpc_lwip_close(socket)
	//	if err != nil {
	//		return err
	//	}
	//	//	fmt.Printf("sending...\r\n")
	//	//	data := []byte{0x00, 0x01, 0x0f, 0x01, byte(seq), byte(seq >> 8), byte(seq >> 16), byte(seq >> 24)}
	//	//	crc := computeCRC16(data)
	//	//
	//	//	buf := []byte{byte(len(data)), byte(len(data) >> 8), byte(crc), byte(crc >> 8)}
	//	//	fmt.Printf("%2d : %#v\n", len(buf), buf)
	//	//	p.Write(buf)
	//	//
	//	//	p.Write(data)
	//	//	fmt.Printf("%2d : %#v\n", len(data), data)
	//	//	fmt.Printf("sending...done\r\n")
	//	//
	//	//	buf = make([]byte, 4)
	//	//	fmt.Printf("reading...\r\n")
	//	//	n, err := io.ReadFull(p, buf)
	//	//	fmt.Printf("reading...done\r\n")
	//	//	fmt.Printf("%2d : %#v\n", n, buf[:n])
	//	//
	//	//	length := uint16(buf[0]) + uint16(buf[1])<<8
	//	//	crc = uint16(buf[2]) + uint16(buf[3])<<8
	//	//
	//	//	fmt.Printf("length : %d\n", length)
	//	//	fmt.Printf("crc16  : 0x%04X\n", crc)
	//	//
	//	//	buf = make([]byte, length)
	//	//	fmt.Printf("reading...\r\n")
	//	//	n, err = io.ReadFull(p, buf)
	//	//	fmt.Printf("reading...done\r\n")
	//	//	fmt.Printf("%2d : %#v\n", n, buf[:n])

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
	fmt.Printf("IP      : %d.%d.%d.%d\n", byte(i.IP>>24), byte(i.IP>>16), byte(i.IP>>8), byte(i.IP))
	fmt.Printf("Mask    : %d.%d.%d.%d\n", byte(i.Mask>>24), byte(i.Mask>>16), byte(i.Mask>>8), byte(i.Mask))
	fmt.Printf("Gateway : %d.%d.%d.%d\n", byte(i.Gateway>>24), byte(i.Gateway>>16), byte(i.Gateway>>8), byte(i.Gateway))
}
