package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.ListenPacket("udp4", ":8899")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s sent: %s\n", addr, buf[:n])

		if buf[0] != '*' {

			broadcastAddr := &net.UDPAddr{
				IP:   net.IPv4(255, 255, 255, 255),
				Port: 8899,
			}

			t := time.Now().Nanosecond()
			n, err = conn.WriteTo(
				[]byte(fmt.Sprintf("*X -> %v -> %v", string(buf[:n]), t)),
				broadcastAddr,
			)
			if err != nil {
				panic(err)
			}

			fmt.Printf("sent back %d bytes to %s\n", n, addr)
		}
	}

}
