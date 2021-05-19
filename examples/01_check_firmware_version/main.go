package main

import (
	"fmt"
	"time"
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

	ver, err := rtl.Rpc_system_version()
	if err != nil {
		return err
	}
	fmt.Printf("RTL8720 Firmware Version: %#v\r\n", ver)

	for {
		time.Sleep(1 * time.Hour)
	}

	return nil
}
