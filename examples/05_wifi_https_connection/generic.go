// +build !baremetal

package main

import (
	"flag"

	"github.com/sago35/go-erpcgen/rtl8720dn"
	"go.bug.st/serial"
)

func setupRTL8720DN() (*rtl8720dn.RTL8720DN, error) {
	adapterInit := flag.Bool("init", false, "init")
	port := flag.String("port", "", "port")
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()

	p, err := serial.Open(*port, &serial.Mode{BaudRate: 115200})
	if err != nil {
		return nil, err
	}

	rtl := rtl8720dn.New(p)
	rtl.Debug(*debug)

	if *adapterInit {
		_, err = rtl.Rpc_tcpip_adapter_init()
		if err != nil {
			return nil, err
		}
	}

	return rtl, nil
}
