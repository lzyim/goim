package client

import (
	"bufio"
	"github.com/soloslee/goim/message"
	"log"
	"net"
)

type ClientMap map[string]*Client

type Client struct {
	Uuid   string
	conn   net.Conn
	To     chan *message.Message
	recive chan string
	Quit   chan *Client
	reader *bufio.Reader
	writer *bufio.Writer
}

func New(uuid string, conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{uuid, conn, make(chan *message.Message), make(chan string), make(chan *Client), reader, writer}
	client.listen()
	return client
}

func (client *Client) read() {
	for {
		if line, _, err := client.reader.ReadLine(); err == nil {
			msg := &message.Message{}
			if err := msg.Unmarshal(line, msg); err == nil {
				client.To <- &message.Message{Fuuid: client.Uuid, Tuuid: msg.Tuuid, Data: msg.Data}
			}
		} else {
			client.Quit <- client
			return
		}
	}
}

func (client *Client) write() {
	for res := range client.recive {
		if _, err := client.writer.WriteString(res + "\n"); err != nil {
			return
		}
		if err := client.writer.Flush(); err != nil {
			log.Printf("Write error: %s\n", err)
			return
		}
	}
}

func (client *Client) listen() {
	go client.read()
	go client.write()
}

func (client *Client) Close() {
	close(client.To)
	close(client.recive)
	client.conn.Close()
}

func (client *Client) Recive(msg string) {
	client.recive <- msg
}
