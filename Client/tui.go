package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type screenClient struct {
	app               *tview.Application
	gridScreen        *tview.Grid
	chatView          *tview.TextView
	onlineUsersArea   *tview.TextView
	commandList       *tview.TextArea
	sendMessagesInput *tview.InputField
	sendMessagesChan  chan string
	receiveMessages   chan string
}

func initModel() screenClient {
	gridScreen := tview.NewGrid()
	gridScreen.SetRows(0, 0, 1).
		SetColumns(0, 40).
		SetBorder(true).
		SetTitle("[Chat 0.3]").
		SetTitleAlign(tview.AlignLeft)

	chatView := tview.NewTextView()
	chatView.SetBorder(true)

	onlineUsersArea := tview.NewTextView()
	onlineUsersArea.SetText("no one online here").
		SetBorder(true).
		SetTitle(" Users ")

	commandsList := tview.NewTextArea()
	commandsList.SetBorder(true).
		SetTitle(" Commands ")

	sendMessagesChan := make(chan string)

	sendMessagesInput := tview.NewInputField()
	sendMessagesInput.SetPlaceholder("Write here: ")

	sendMessagesInput.SetDoneFunc(func(key tcell.Key) {
		sendMessagesChan <- fmt.Sprintf("%s\n", sendMessagesInput.GetText())
		sendMessagesInput.SetText("")
	})

	gridScreen.AddItem(chatView, 0, 0, 2, 1, 1, 1, false)
	gridScreen.AddItem(onlineUsersArea, 0, 1, 1, 1, 1, 1, false)
	gridScreen.AddItem(commandsList, 1, 1, 1, 1, 1, 1, false)
	gridScreen.AddItem(sendMessagesInput, 2, 0, 1, 1, 1, 1, true)

	app := tview.NewApplication().
		SetRoot(gridScreen, true)
	return screenClient{
		app:               app,
		gridScreen:        gridScreen,
		chatView:          chatView,
		onlineUsersArea:   onlineUsersArea,
		commandList:       commandsList,
		sendMessagesInput: sendMessagesInput,
		sendMessagesChan:  sendMessagesChan,
		receiveMessages:   make(chan string),
	}
}

func (s *screenClient) runScreen() {
	if err := s.app.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *screenClient) updateScreen() {
	for {
		msg := <-s.receiveMessages
		s.chatView.Write([]byte(msg))
		s.app.Draw()
	}
}
