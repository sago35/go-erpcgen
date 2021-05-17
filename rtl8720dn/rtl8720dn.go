package rtl8720dn

import "io"

type RTL8720DN struct {
	port io.ReadWriter
	seq  uint64
}

func New(r io.ReadWriter) *RTL8720DN {
	return &RTL8720DN{
		port: r,
		seq:  1,
	}
}

func (r *RTL8720DN) SetSeq(s uint64) {
	r.seq = s
}
