package server

import (
	"bytes"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/soloslee/goim/client"
	"github.com/soloslee/goim/config"
	"github.com/soloslee/goim/message"
	"log"
	"net"
)

type Server struct {
	port        int
	listener    net.Listener
	client      chan *client.Client
	quit        chan *client.Client
	clientTable client.ClientMap
	message     chan *message.Message
}

// Start start a goim server instance
func Start(config *config.Config) {
	log.Println("Starting goim")
	server := &Server{
		port:        config.Port,
		client:      make(chan *client.Client),
		quit:        make(chan *client.Client),
		clientTable: make(client.ClientMap),
		message:     make(chan *message.Message),
	}
	server.listen()
	server.startTcp()
}

func (server *Server) listen() {
	go func() {
		for {
			select {
			case c := <-server.client:
				server.handleConn(c)
			case m := <-server.message:
				server.handleMsg(m)
			case q := <-server.quit:
				server.handleQuit(q)
			}
		}
	}()
}

func (server *Server) handleConn(client *client.Client) {
	go func() {
		for msg := range client.To {
			if msg == nil {
				return
			}
			server.message <- msg

		}
	}()
	go func() {
		for cli := range client.Quit {
			server.quit <- cli
			return
		}
	}()
}

func (server *Server) handleMsg(msg *message.Message) {
	cli := server.clientTable[msg.Tuuid]
	if cli != nil {
		var buffer bytes.Buffer
		buffer.WriteString("\nUser ")
		buffer.WriteString(msg.Fuuid)
		buffer.WriteString(" send a message to you: ")
		buffer.WriteString(msg.Data)
		buffer.WriteString("\nEnter your friend's uuid: ")
		cli.Recive(buffer.String())
	}
}

func (server *Server) handleQuit(client *client.Client) {
	client.Close()
	delete(server.clientTable, client.Uuid)
}

func (server *Server) startTcp() {
	addr := fmt.Sprintf("127.0.0.1:%d", server.port)
	server.listener, _ = net.Listen("tcp", addr)
	log.Printf("listen on port %d", server.port)
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			log.Fatalln(err)
			return
		}
		cli := client.New(uuid.NewV4().String(), conn)
		server.clientTable[cli.Uuid] = cli
		server.client <- cli
		cli.Recive("\nYour uuid is " + cli.Uuid)
	}
}

// Stop stop a goim server instance
func (server *Server) Stop() {
	server.listener.Close()
	close(server.client)
	close(server.message)
}
