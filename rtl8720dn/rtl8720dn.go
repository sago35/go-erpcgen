package rtl8720dn

import "io"

type RTL8720DN struct {
	port     io.ReadWriter
	seq      uint64
	received chan bool
	debug    bool
}

func New(r io.ReadWriter) *RTL8720DN {
	ret := &RTL8720DN{
		port:     r,
		seq:      1,
		received: make(chan bool, 1),
		debug:    false,
	}

	go ret.readThread()

	return ret
}

func (r *RTL8720DN) SetSeq(s uint64) {
	r.seq = s
}

func (r *RTL8720DN) Debug(b bool) {
	r.debug = b
}
