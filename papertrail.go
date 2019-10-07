package log

import "net"

// Papertrail holds conn data for papertrail
type Papertrail struct {
	Address string
}

// Write implements papertrail io writer
func (p *Papertrail) Write(b []byte) (int, error) {
	var con net.Conn
	con, err := net.Dial("udp", p.Address)
	if err != nil {
		return 0, nil
	}
	return con.Write(b)
}
