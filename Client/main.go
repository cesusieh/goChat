package main

func main() {
	screen := initModel()
	go screen.runScreen()
	go screen.updateScreen()

	conn := connectServer(screen.sendMessagesChan, screen.receiveMessages)
	go conn.listenServer()
	conn.writeOnServer()
}
