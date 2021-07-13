package main

import (
	"fmt"
	"time"
)

var (
	ssid     string
	password string

	server       = `www.example.com`
	test_root_ca = `-----BEGIN CERTIFICATE-----
MIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh
MQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3
d3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD
QTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT
MRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j
b20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB
CSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97
nh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt
43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P
T19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4
gdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO
BgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR
TLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw
DQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr
hMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg
06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF
PnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls
YSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk
CAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=
-----END CERTIFICATE-----
`
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

	client, err := rtl.Rpc_wifi_ssl_client_create()
	if err != nil {
		return err
	}

	err = rtl.Rpc_wifi_ssl_init(client)
	if err != nil {
		return err
	}

	err = rtl.Rpc_wifi_ssl_set_timeout(client, 0x0001D4C0)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_wifi_ssl_set_rootCA(client, test_root_ca)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_wifi_start_ssl_client(client, server, 443, 0x0001D4C0)
	if err != nil {
		return err
	}

	_, err = rtl.Rpc_wifi_ssl_get_socket(client)
	if err != nil {
		return err
	}

	fmt.Printf("Connected to server!\r\n")

	req := "GET https://" + server + " HTTP/1.0\r\nHost: " + server + "\r\nConnection: close\r\n\r\n"
	_, err = rtl.Rpc_wifi_send_ssl_data(client, []byte(req), uint16(len(req)))
	if err != nil {
		return err
	}

	// rpc_lwip_recv()
	// rpc_lwip_errno()

	rcvBuf := make([]byte, 0x0400)

	for {
		n, err := rtl.Rpc_wifi_get_ssl_receive(client, &rcvBuf, int32(cap(rcvBuf)))
		if err != nil {
			return err
		}
		if n == -76 {
			time.Sleep(500 * time.Millisecond)
		} else if n == 0 || n == -30848 {
			break
		} else if n < 0 {
			return fmt.Errorf("error %d", n)
		} else {
			fmt.Printf("%s\r\n", rcvBuf)
		}
	}

	err = rtl.Rpc_wifi_stop_ssl_socket(client)
	if err != nil {
		return err
	}

	err = rtl.Rpc_wifi_ssl_client_destroy(client)
	if err != nil {
		return err
	}

	for {
		time.Sleep(1 * time.Hour)
	}

	return nil
}
