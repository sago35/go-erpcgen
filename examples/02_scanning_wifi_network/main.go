package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
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

	for {
		fmt.Printf("scan start\n")
		_, err := rtl.Rpc_wifi_scan_start()
		if err != nil {
			return err
		}
		time.Sleep(5000 * time.Millisecond)

		num := uint16(0)
		for num == 0 {
			num, err = rtl.Rpc_wifi_scan_get_ap_num()
			if err != nil {
				return err
			}
			time.Sleep(1000 * time.Millisecond)
		}

		fmt.Printf("%d networks found\n", num)

		result := make([]byte, 66*num)
		_, err = rtl.Rpc_wifi_scan_get_ap_records(num, result[:])
		if err != nil {
			return err
		}

		records := parseWifiApRecord(int(num), result)
		for i, r := range records {
			fmt.Printf("%d: %s (%d)*\n", i, r.SSID, r.RSSI)
		}

		time.Sleep(5000 * time.Millisecond)
	}

	return nil
}

type WifiApRecord struct {
	SSID string
	RSSI int // signal strength of AP
}

func parseWifiApRecord(num int, buf []byte) []WifiApRecord {
	ret := []WifiApRecord{}

	for i := 0; i < num; i++ {
		start := i*62 + 1
		end := 0
		for j, b := range buf[start : start+62] {
			if b != 0 {
				end = j
			} else {
				break
			}
		}
		ret = append(ret, WifiApRecord{
			SSID: string(buf[start : start+end]),
			RSSI: -1 * (256 - int(buf[start+39])),
		})
	}
	return ret
}
