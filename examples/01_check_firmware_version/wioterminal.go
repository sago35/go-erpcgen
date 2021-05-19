// +build wioterminal

package main

import (
	"machine"
	"time"

	"github.com/sago35/go-erpcgen/rtl8720dn"
)

func setupRTL8720DN() (*rtl8720dn.RTL8720DN, error) {
	debug := true

	waitSerial()

	uart2 := machine.UART2
	uart2.Configure(machine.UARTConfig{TX: machine.UART2_TX_PIN, RX: machine.UART2_RX_PIN, BaudRate: 614400})

	machine.RTL8720D_CHIP_PU.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.RTL8720D_CHIP_PU.Low()
	time.Sleep(100 * time.Millisecond)
	machine.RTL8720D_CHIP_PU.High()
	time.Sleep(100 * time.Millisecond)

	rtl := rtl8720dn.New(uart2)
	rtl.Debug(debug)

	_, err := rtl.Rpc_tcpip_adapter_init()
	if err != nil {
		return nil, err
	}

	return rtl, nil
}

// Wait for user to open serial console
func waitSerial() {
	for !machine.Serial.DTR() {
		time.Sleep(100 * time.Millisecond)
	}
}
