package main

import (
	"fmt"
	"log"
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

	ver, err := rtl.Rpc_system_version()
	if err != nil {
		return err
	}
	fmt.Printf("RTL8720 Firmware Version: %#v\n", ver)

	return nil
}
