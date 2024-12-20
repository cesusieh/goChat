package main

import (
	"bufio"
	"fmt"
	"net"
)

type server struct {
	listener        net.Listener
	clientList      []client
	receiveMessages chan string
	onlineUserChan  chan client
}

type client struct {
	conn     net.Conn
	nickname string
}

func initServer(receiveMessages chan string, onlineUserChan chan client) server {
	receiveMessages <- "Iniciando servidor. . .\n"

	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		receiveMessages <- fmt.Sprintf("Impossível iniciar o servidor.\n%v", err)
	}
	receiveMessages <- "Servidor iniciado com sucesso!\n\n"
	return server{
		listener:        listener,
		clientList:      []client{},
		receiveMessages: receiveMessages,
		onlineUserChan:  onlineUserChan,
	}
}

func (s *server) manageServer() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.receiveMessages <- fmt.Sprintf("Erro ao receber conexão.\n%v", err)
		}

		nick := requestNick(conn)

		client := client{
			conn:     conn,
			nickname: nick,
		}

		s.receiveMessages <- fmt.Sprintf("[%s] - Conexão recebida.\n", client.nickname)
		s.onlineUserChan <- client

		s.clientList = append(s.clientList, client)
		go client.listenAndRepass(&s.clientList)
	}
}

func (c client) listenAndRepass(clientList *[]client) {
	reader := bufio.NewReader(c.conn)
	for {
		msg, _ := reader.ReadString('\n')
		for _, client := range *clientList {
			if c != client {
				client.conn.Write([]byte(fmt.Sprintf("[%s]:%s", c.nickname, msg)))
			}
		}
	}
}

func requestNick(conn net.Conn) string {
	conn.Write([]byte("Informe seu nick.\n"))

	reader := bufio.NewReader(conn)
	nick, _ := reader.ReadString('\n')
	nick = nick[:len(nick)-1]

	conn.Write([]byte("Seja bem vindo :)\n\n"))
	return nick
}
