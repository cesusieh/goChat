package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type connServer struct {
	conn            net.Conn
	receiveMessages chan string
	sendMessages    chan string
}

func connectServer(sendMessages, receiveMessages chan string) connServer {
	receiveMessages <- "Tentando se conectar ao servidor. . .\n"
	conn, err := net.Dial("tcp", "localhost:9001")
	if err != nil {
		receiveMessages <- fmt.Sprintf("Impossível se conectar ao servidor.\n%v", err)
	}
	receiveMessages <- "Conectado com sucesso!\n\n"
	return connServer{
		conn:            conn,
		sendMessages:    sendMessages,
		receiveMessages: receiveMessages,
	}
}

func (s connServer) listenServer() {
	reader := bufio.NewReader(s.conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s.receiveMessages <- msg
	}
}

func (s connServer) writeOnServer() {
	for {
		msg := <-s.sendMessages
		s.receiveMessages <- fmt.Sprintf("[Eu]:%s", msg)
		_, err := s.conn.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	}
}
