package rtl8720dn

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func (r *RTL8720DN) ConnectToAP(ssid string, password string) error {
	if len(ssid) == 0 || len(password) == 0 {
		return fmt.Errorf("connection failed: either ssid or password not set")
	}

	_, err := r.Rpc_wifi_off()
	if err != nil {
		return err
	}
	_, err = r.Rpc_wifi_on(0x00000001)
	if err != nil {
		return err
	}

	_, err = r.Rpc_wifi_disconnect()
	if err != nil {
		return err
	}

	numTry := 5
	securityType := uint32(0x00400004)
	for i := 0; i < numTry; i++ {
		ret, err := r.Rpc_wifi_connect(ssid, password, securityType, -1, 0)
		if err != nil {
			return err
		}
		if ret != 0 {
			if i == numTry-1 {
				return fmt.Errorf("connection failed: rpc_wifi_connect failed")
			}
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	_, err = r.Rpc_tcpip_adapter_dhcpc_start(0)
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		_, err = r.Rpc_wifi_is_connected_to_ap()
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (r *RTL8720DN) GetIP() (ip, subnet, gateway IPAddress, err error) {
	ip_info := make([]byte, 12)
	_, err = r.Rpc_tcpip_adapter_get_ip_info(0, ip_info)
	if err != nil {
		return nil, nil, nil, err
	}

	ip = IPAddress(ip_info[0:4])
	subnet = IPAddress(ip_info[4:8])
	gateway = IPAddress(ip_info[8:12])

	return ip, subnet, gateway, nil
}

type IPAddress []byte

func (addr IPAddress) String() string {
	if len(addr) < 4 {
		return ""
	}
	return strconv.Itoa(int(addr[0])) + "." + strconv.Itoa(int(addr[1])) + "." + strconv.Itoa(int(addr[2])) + "." + strconv.Itoa(int(addr[3]))
}

func ParseIPv4(s string) (IPAddress, error) {
	v := strings.Split(s, ".")
	v0, _ := strconv.Atoi(v[0])
	v1, _ := strconv.Atoi(v[1])
	v2, _ := strconv.Atoi(v[2])
	v3, _ := strconv.Atoi(v[3])
	return IPAddress([]byte{byte(v0), byte(v1), byte(v2), byte(v3)}), nil
}

func (addr IPAddress) AsUint32() uint32 {
	if len(addr) < 4 {
		return 0
	}
	b := []byte(string(addr))
	return binary.BigEndian.Uint32(b[0:4])
}
