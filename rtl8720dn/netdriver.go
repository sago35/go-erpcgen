package rtl8720dn

import (
	"fmt"
	"strconv"
)

// Here is the implementation of tinygo-org/x/drivers/net.DeviceDriver.

func (r *RTL8720DN) GetDNS(domain string) (string, error) {
	if r.debug {
		fmt.Printf("GetDNS(%q)\r\n", domain)
	}

	ipaddr := [4]byte{}
	_, err := r.Rpc_netconn_gethostbyname(domain, ipaddr[:])
	if err != nil {
		return "", err
	}

	ret, err := fmt.Sprintf("%d.%d.%d.%d", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]), nil
	if r.debug {
		fmt.Printf("-> %s\r\n", ret)
		fmt.Printf("-> %02X.%02X.%02X.%02X\r\n", ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3])
	}
	return ret, err
}

func (r *RTL8720DN) ConnectTCPSocket(addr, port string) error {
	if r.debug {
		fmt.Printf("ConnectTCPSocket(%q, %q)\r\n", addr, port)
	}

	ipaddr := [4]byte{}
	_, err := r.Rpc_netconn_gethostbyname(addr, ipaddr[:])
	if err != nil {
		return err
	}

	portNum, err := strconv.ParseUint(port, 0, 0)
	if err != nil {
		return err
	}

	socket, err := r.Rpc_lwip_socket(0x02, 0x01, 0x00)
	if err != nil {
		return err
	}
	r.socket = socket
	r.connectionType = ConnectionTypeTCP

	_, err = r.Rpc_lwip_fcntl(socket, 0x00000003, 0x00000000)
	if err != nil {
		return err
	}

	_, err = r.Rpc_lwip_fcntl(socket, 0x00000004, 0x00000001)
	if err != nil {
		return err
	}

	name := []byte{0x00, 0x02, 0x00, 0x50, 0xC0, 0xA8, 0x01, 0x76, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	name[2] = byte(portNum >> 8)
	name[3] = byte(portNum)
	name[4] = byte(ipaddr[0])
	name[5] = byte(ipaddr[1])
	name[6] = byte(ipaddr[2])
	name[7] = byte(ipaddr[3])

	_, err = r.Rpc_lwip_connect(socket, name, uint32(len(name)))
	if err != nil {
		return err
	}

	readset := []byte{}
	writeset := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	exceptset := []byte{}
	timeout := []byte{}
	_, err = r.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	if err != nil {
		return err
	}

	optlen := uint32(4)
	_, err = r.Rpc_lwip_getsockopt(socket, 0x00000FFF, 0x00001007, []byte{0xA5, 0xA5, 0xA5, 0xA5}, nil, optlen)
	if err != nil {
		return err
	}

	_, err = r.Rpc_lwip_fcntl(socket, 0x00000003, 0x00000000)
	if err != nil {
		return err
	}

	_, err = r.Rpc_lwip_fcntl(socket, 0x00000004, 0x00000000)
	if err != nil {
		return err
	}

	readset = []byte{}
	writeset = []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	exceptset = []byte{}
	timeout = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x40, 0x42, 0x0F, 0x00, 0xFF, 0xFF, 0xFF, 0xFF}
	_, err = r.Rpc_lwip_select(0x01, readset, writeset, exceptset, timeout)
	if err != nil {
		return err
	}

	return nil
}

func (r *RTL8720DN) ConnectSSLSocket(addr, port string) error {
	if r.debug {
		fmt.Printf("ConnectSSLSocket(%q, %q)\r\n", addr, port)
	}

	client, err := r.Rpc_wifi_ssl_client_create()
	if err != nil {
		return err
	}
	r.client = client
	r.connectionType = ConnectionTypeTLS

	err = r.Rpc_wifi_ssl_init(client)
	if err != nil {
		return err
	}

	err = r.Rpc_wifi_ssl_set_timeout(client, 0x0001D4C0)
	if err != nil {
		return err
	}

	_, err = r.Rpc_wifi_ssl_set_rootCA(client, *r.root_ca)
	if err != nil {
		return err
	}

	// TODO: use port
	_, err = r.Rpc_wifi_start_ssl_client(client, addr, 443, 0x0001D4C0)
	if err != nil {
		return err
	}

	_, err = r.Rpc_wifi_ssl_get_socket(client)
	if err != nil {
		return err
	}
	return nil
}

func (r *RTL8720DN) ConnectUDPSocket(addr, sendport, listenport string) error {
	if r.debug {
		fmt.Printf("ConnectUDPSocket(%q, %q, %q)\r\n", addr, sendport, listenport)
	}
	fmt.Printf("not implemented yet\r\n")
	return nil
}

func (r *RTL8720DN) DisconnectSocket() error {
	if r.debug {
		fmt.Printf("DisconnectSocket()\r\n")
	}
	switch r.connectionType {
	case ConnectionTypeTCP:
		_, err := r.Rpc_lwip_close(r.socket)
		if err != nil {
			return err
		}
	case ConnectionTypeTLS:
		err := r.Rpc_wifi_stop_ssl_socket(r.client)
		if err != nil {
			return err
		}

		err = r.Rpc_wifi_ssl_client_destroy(r.client)
		if err != nil {
			return err
		}
	default:
	}
	r.connectionType = ConnectionTypeNone
	return nil
}

func (r *RTL8720DN) StartSocketSend(size int) error {
	if r.debug {
		fmt.Printf("StartSocketSend(%d)\r\n", size)
	}
	// No implementation required
	return nil
}

func (r *RTL8720DN) Write(b []byte) (n int, err error) {
	if r.debug {
		fmt.Printf("Write(%#v)\r\n", b)
	}

	switch r.connectionType {
	case ConnectionTypeTCP:
		sn, err := r.Rpc_lwip_send(r.socket, b, 0x00000008)
		if err != nil {
			return 0, err
		}
		n = int(sn)
	case ConnectionTypeTLS:
		sn, err := r.Rpc_wifi_send_ssl_data(r.client, b, uint16(len(b)))
		if err != nil {
			return 0, err
		}
		n = int(sn)
	default:
		return 0, nil
	}
	return n, nil
}

func (r *RTL8720DN) ReadSocket(b []byte) (n int, err error) {
	if r.debug {
		//fmt.Printf("ReadSocket(b)\r\n")
	}
	if r.connectionType == ConnectionTypeNone {
		return 0, nil
	}

	switch r.connectionType {
	case ConnectionTypeTCP:
		nn, err := r.Rpc_lwip_recv(r.socket, b, uint32(len(b)), 0x00000008, 0x00002800)
		if err != nil {
			return 0, err
		}

		if nn == -1 {
			return 0, nil
		} else if nn == 0 {
			return 0, r.DisconnectSocket()
		}
		if r.length == 0 {
			header := httpHeader(b[:nn])
			r.length = header.ContentLength()
		}
		r.length -= int(nn)
		if r.length == 0 {
			return int(nn), r.DisconnectSocket()
		}
		n = int(nn)
	case ConnectionTypeTLS:
		nn, err := r.Rpc_wifi_get_ssl_receive(r.client, b, int32(len(b)))
		if err != nil {
			return 0, err
		}
		if nn < 0 {
			return 0, fmt.Errorf("error %d", n)
		} else if nn == 0 || nn == -30848 {
			return 0, r.DisconnectSocket()
		}
		n = int(nn)
	default:
	}

	return n, nil
}

func (r *RTL8720DN) IsSocketDataAvailable() bool {
	if r.debug {
		fmt.Printf("IsSocketDataAvailable()\r\n")
	}
	fmt.Printf("not implemented yet\r\n")
	return true
}

func (r *RTL8720DN) Response(timeout int) ([]byte, error) {
	if r.debug {
		fmt.Printf("Response(%d))\r\n", timeout)
	}
	// No implementation required
	return nil, nil
}
