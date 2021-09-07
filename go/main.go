package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func main() {

	quit := make(chan bool)

	go listen()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'quit' to stop reading UDP messages:\r\n")
	text, _ := reader.ReadString('\n')
	if text == "quit" {
		close(quit)
	}
}

func listen() {
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
