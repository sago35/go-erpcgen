package rtl8720dn

import (
	"fmt"
)

var (
	headerBuf            [4]byte
	readBuf              [4]byte
	startWriteMessageBuf [1024]byte
)

const (
	xVersion = 1
)

func startWriteMessage(msgType, service, requestNumber, sequence uint32) []byte {
	startWriteMessageBuf[0] = byte(msgType)
	startWriteMessageBuf[1] = byte(requestNumber)
	startWriteMessageBuf[2] = byte(service)
	startWriteMessageBuf[3] = byte(xVersion)

	startWriteMessageBuf[4] = byte(sequence)
	startWriteMessageBuf[5] = byte(sequence >> 8)
	startWriteMessageBuf[6] = byte(sequence >> 16)
	startWriteMessageBuf[7] = byte(sequence >> 24)

	return startWriteMessageBuf[:8]
}

func (r *RTL8720DN) performRequest(msg []byte) error {
	crc := computeCRC16(msg)
	headerBuf[0] = byte(len(msg))
	headerBuf[1] = byte(len(msg) >> 8)
	headerBuf[2] = byte(crc)
	headerBuf[3] = byte(crc >> 8)

	if r.debug {
		fmt.Printf("tx : %2d : %s\n", len(headerBuf), dumpHex(headerBuf[:]))
	}

	r.port.Write(headerBuf[:])

	if r.debug {
		fmt.Printf("tx : %2d : %s\n", len(msg), dumpHex(msg))
	}
	r.port.Write(msg)

	//n, err := io.ReadFull(r.port, headerBuf[:])
	//if err != nil {
	//	return err
	//}
	//if Debug {
	//	fmt.Printf("rx : %2d : %#v\n", n, headerBuf[:n])
	//}

	//n, err = io.ReadFull(r.port, startWriteMessageBuf[:8])
	//if err != nil {
	//	return err
	//}
	//if Debug {
	//	fmt.Printf("rx : %2d : %#v\n", n, startWriteMessageBuf[:n])
	//}
	return nil
}