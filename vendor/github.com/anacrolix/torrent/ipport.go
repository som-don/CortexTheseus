package torrent

import (
	"net"
	"strconv"
)

// Extracts the port as an integer from an address string.
func addrPortOrZero(addr net.Addr) int {
	switch raw := addr.(type) {
	case *net.UDPAddr:
		return raw.Port
	case *net.TCPAddr:
		return raw.Port
	default:
		_, port, err := net.SplitHostPort(addr.String())
		if err != nil {
			return 0
		}
		i64, err := strconv.ParseInt(port, 0, 0)
		if err != nil {
			panic(err)
		}
		return int(i64)
	}
}

func addrIpOrNil(addr net.Addr) net.IP {
	if addr == nil {
		return nil
	}
	switch raw := addr.(type) {
	case *net.UDPAddr:
		return raw.IP
	case *net.TCPAddr:
		return raw.IP
	default:
		host, _, err := net.SplitHostPort(addr.String())
		if err != nil {
			return nil
		}
		return net.ParseIP(host)
	}
}

type ipPortAddr struct {
	IP   net.IP
	Port int
}

func (ipPortAddr) Network() string {
	return ""
}

func (me ipPortAddr) String() string {
	return net.JoinHostPort(me.IP.String(), strconv.FormatInt(int64(me.Port), 10))
}

func tryIpPortFromNetAddr(na net.Addr) (ret ipPortAddr, ok bool) {
	ret.IP = addrIpOrNil(na)
	if ret.IP == nil {
		return
	}
	ret.Port = addrPortOrZero(na)
	ok = true
	return
}
