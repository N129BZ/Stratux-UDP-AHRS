package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func main() {

	p := make([]byte, 512)
	addr := net.UDPAddr{
		Port: 4000,
		IP:   net.ParseIP("0.0.0.0"),
	}

	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	defer conn.Close()

	for {
		conn.ReadFromUDP(p)
		if p[1] == 76 {
			ahrs := NewAhrsRecord()
			err = ahrs.Read(kaitai.NewStream(bytes.NewReader(p)), ahrs)
			if err != nil {
				fmt.Printf("Error  %v", err)
				continue
			}

			b, err := json.Marshal(ahrs)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(string(string(b)))
		}
	}
}
