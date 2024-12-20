package main

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

type screenServer struct {
	app             *tview.Application
	gridScreen      *tview.Grid
	chatView        *tview.TextView
	onlineUsersArea *tview.TextView
	hehe            *tview.TextView
	receiveMessages chan string
	onlineUserChan  chan client
}

func initModel() screenServer {
	grid := tview.NewGrid()
	grid.SetRows(0, 0).
		SetColumns(0, 0).
		SetBorder(true).
		SetTitle("[Chat 0.3]").
		SetTitleAlign(tview.AlignLeft)

	chatView := tview.NewTextView()
	chatView.SetBorder(true)

	onlineUsersAreas := tview.NewTextView()
	onlineUsersAreas.SetBorder(true).
		SetTitle(" Users ")

	hehe := tview.NewTextView()
	hehe.SetBorder(true).
		SetTitle(" Hehe ")

	grid.AddItem(chatView, 0, 0, 2, 1, 1, 1, false)
	grid.AddItem(onlineUsersAreas, 0, 1, 1, 1, 1, 1, false)
	grid.AddItem(hehe, 1, 1, 1, 1, 1, 1, false)

	app := tview.NewApplication().SetRoot(grid, true)
	return screenServer{
		app:             app,
		gridScreen:      grid,
		chatView:        chatView,
		onlineUsersArea: onlineUsersAreas,
		hehe:            hehe,
		receiveMessages: make(chan string),
		onlineUserChan:  make(chan client),
	}
}

func (s *screenServer) runScreen() {
	if err := s.app.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *screenServer) listenMsgChan() {
	for {
		msg := <-s.receiveMessages
		s.chatView.Write([]byte(msg))
		s.app.Draw()
	}
}

func (s *screenServer) receiveUser() {
	for {
		client := <-s.onlineUserChan
		s.onlineUsersArea.Write([]byte(fmt.Sprintf("%s - %s\n", client.nickname, client.conn.LocalAddr())))
		s.app.Draw()
	}
}
