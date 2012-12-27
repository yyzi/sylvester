package connections

import (
	"errors"
	"log"
	"net"
)

var (
	ErrTCPConnection = errors.New("[err] TCP resolution or connection error")
	ErrUDPConnection = errors.New("[err] UDP resolution or connection error")
)

func TcpConnect(address string) (c *net.TCPConn, err error) {
	log.Printf("[TCP] Dialing %s", address)
	tcpaddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, ErrTCPConnection
	}

	c, err = net.DialTCP("tcp", nil, tcpaddr)
	if err != nil {
		return nil, ErrTCPConnection
	}

	c.SetKeepAlive(true)

	return c, nil
}

func UdpConnect(address string) (c *net.UDPConn, err error) {
	log.Printf("[UDP] Dialing %s", address)
	udpaddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return nil, ErrUDPConnection
	}

	c, err = net.DialUDP("udp", nil, udpaddr)
	if err != nil {
		return nil, ErrUDPConnection
	}

	return c, nil
}