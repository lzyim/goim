package server

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/soloslee/goim/message"
	"net"
	"testing"
)

func TestConn(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	reader := bufio.NewReader(conn)
	if line, _, err := reader.ReadLine(); err == nil {
		if _, err := uuid.FromString(string(line)); err == nil {
			conn.Close()
		} else {
			t.Error("Expected no error, but there was one: %v", err)
		}
	}
}

func TestSendMsg(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	uuidStr, _, _ := reader.ReadLine()
	uuid, _ := uuid.FromString(string(uuidStr))
	data := []byte("Hello, world!")
	msg := &message.Message{Tuuid: uuid, Data: data}
	ok := make(chan bool)
	res, err := msg.Marshal()
	go func() {
		for {
			if line, _, err := reader.ReadLine(); err == nil {
				rec := fmt.Sprintf("User %s send a message to you: Hello, world!", uuid)
				if bytes.Equal([]byte(rec), line) {
					ok <- true
					conn.Close()
					return
				} else {
					t.Error("Expected no error, but there was one: %v", err)
				}
			} else {
				t.Error("Expected no error, but there was one: %v", err)
			}
		}
	}()
	if err == nil {
		writer.Write(append(res, []byte("\n")...))
		writer.Flush()
	}
	<-ok
}
