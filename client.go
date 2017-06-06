package main

import (
	"bufio"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/soloslee/goim/message"
	"net"
)

func getData(conn *net.Conn) {
	for {
		var uuidStr string
		var data []byte
		fmt.Print("Enter your friend's uuid: ")
		fmt.Scanln(&uuidStr)
		uuid, _ := uuid.FromString(uuidStr)
		fmt.Print("Enter a message: ")
		fmt.Scanln(&data)
		msg := &message.Message{Tuuid: uuid, Data: data}
		res, err := msg.Marshal()
		if err == nil {
			fmt.Fprintf(*conn, string(res)+"\n")
		}
	}
}

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	reader := bufio.NewReader(conn)
	go func() {
		for {
			if line, _, err := reader.ReadLine(); err == nil {
				fmt.Println(string(line))
			} else {
				return
			}
		}
	}()
	getData(&conn)
}
