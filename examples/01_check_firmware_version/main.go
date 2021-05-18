package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sago35/go-erpcgen/rtl8720dn"
	"go.bug.st/serial"
)

func main() {
	port := flag.String("port", "", "port")
	debug := flag.Bool("debug", false, "debug")
	flag.Parse()

	err := run(*port, *debug)
	if err != nil {
		log.Fatal(err)
	}
}

func run(port string, debug bool) error {
	p, err := serial.Open(port, &serial.Mode{BaudRate: 115200})
	if err != nil {
		return err
	}

	rtl := rtl8720dn.New(p)
	rtl.Debug(debug)

	ver, err := rtl.Rpc_system_version()
	if err != nil {
		return err
	}
	fmt.Printf("RTL8720 Firmware Version: %#v\n", ver)

	return nil
}
