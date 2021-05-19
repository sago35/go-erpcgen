// +build wioterminal

package main

import (
	"device/sam"
	"machine"
	"runtime/interrupt"
	"time"

	"github.com/sago35/go-erpcgen/rtl8720dn"
)

var (
	uart *machine.UART
)

func handleInterrupt(interrupt.Interrupt) {
	// should reset IRQ
	uart.Receive(byte((uart.Bus.DATA.Get() & 0xFF)))
	uart.Bus.INTFLAG.SetBits(sam.SERCOM_USART_INT_INTFLAG_RXC)
}

func setupRTL8720DN() (*rtl8720dn.RTL8720DN, error) {
	debug := true

	uart = &machine.UART{
		Buffer: machine.NewRingBuffer(),
		Bus:    sam.SERCOM0_USART_INT,
		SERCOM: 0,
	}

	uart.Interrupt = interrupt.New(sam.IRQ_SERCOM0_2, handleInterrupt)
	uart.Configure(machine.UARTConfig{TX: machine.PB24, RX: machine.PC24, BaudRate: 614400})

	machine.RTL8720D_CHIP_PU.Configure(machine.PinConfig{Mode: machine.PinOutput})
	machine.RTL8720D_CHIP_PU.Low()
	time.Sleep(100 * time.Millisecond)
	machine.RTL8720D_CHIP_PU.High()
	time.Sleep(1000 * time.Millisecond)
	waitSerial()

	rtl := rtl8720dn.New(uart)
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
