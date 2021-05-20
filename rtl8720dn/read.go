package rtl8720dn

import (
	"fmt"
	"io"
	"time"
)

func dumpHex(b []byte) {
	for i := range b {
		if i == 0 {
			fmt.Printf("%02X", b[i])
		} else {
			fmt.Printf(" %02X", b[i])
		}
	}
}

var (
	buf     [4]byte
	payload [1024]byte
)

func (r *RTL8720DN) readThread() {
	for {
		n, _ := io.ReadFull(r.port, buf[:4])
		if n == 0 {
			continue
		}

		if r.debug {
			fmt.Printf("rx : %2d : ", n)
			dumpHex(buf[:n])
			fmt.Printf("\r\n")
		}

		length := uint16(buf[0]) + uint16(buf[1])<<8
		crc := uint16(buf[2]) + uint16(buf[3])<<8
		//fmt.Printf("len %d (%X) crc %04X\n", length, length, crc)

		n, _ = io.ReadFull(r.port, payload[:length])
		if r.debug {
			fmt.Printf("rx : %2d : ", length)
			dumpHex(payload[0:n])
			fmt.Printf("\r\n")
		}

		n = int(length)
		//fmt.Printf("rx : %2d : %s\n", n, dumpHex(payload[:n]))

		crcNew := computeCRC16(payload[:n])
		if g, e := crcNew, crc; g != e {
			fmt.Printf("err CRC16: got %04X want %04X\n", g, e)
		}
		if payload[0] == 0x02 {
			r.received <- true

			// switch goroutine
			time.Sleep(1)
		}
	}
}
