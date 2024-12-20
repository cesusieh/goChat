package main

func main() {
	screen := initModel()
	go screen.runScreen()
	go screen.receiveUser()
	go screen.listenMsgChan()

	server := initServer(screen.receiveMessages, screen.onlineUserChan)
	server.manageServer()
}
