package rtl8720dn

import (
	"fmt"
	"io"
)

func dumpHex(b []byte) string {
	ret := ``
	for i := range b {
		if i == 0 {
			ret += fmt.Sprintf("%02X", b[i])
		} else {
			ret += fmt.Sprintf(" %02X", b[i])
		}
	}
	return ret
}

var (
	buf     [4]byte
	payload [1024]byte
)

func ReadThread(r io.ReadWriter) {
	p = r
	received = make(chan bool, 1)
	for {
		n, _ := io.ReadFull(p, buf[:])
		fmt.Printf("rx : %2d : %s\n", n, dumpHex(buf[:n]))

		length := uint16(buf[0]) + uint16(buf[1])<<8
		crc := uint16(buf[2]) + uint16(buf[3])<<8
		//fmt.Printf("len %d (%X) crc %04X\n", length, length, crc)

		//n, _ = io.ReadFull(p, payload[:length])
		fmt.Printf("rx : %2d : ", length)
		for i := 0; i < int(length); i++ {
			n, _ = io.ReadFull(p, payload[i:i+1])
			if i == 0 {
				fmt.Printf("%02X", payload[i:i+1])
			} else {
				fmt.Printf(" %02X", payload[i:i+1])
			}
		}
		fmt.Printf("\n")
		n = int(length)
		//fmt.Printf("rx : %2d : %s\n", n, dumpHex(payload[:n]))

		crcNew := computeCRC16(payload[:n])
		if g, e := crcNew, crc; g != e {
			fmt.Printf("err CRC16: got %04X want %04X\n", g, e)
		}
		if payload[0] == 0x02 {
			received <- true
		}
	}
}
