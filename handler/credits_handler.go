package handler

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

var (
	creditsInstructions = "Credits\n" +
		"======\n" +
		"Bot Profile Picture -> Chelsey (CH21-1)"
)

type CreditsHandler struct{}

func (handler CreditsHandler) Handle(bot *linebot.Client, event *linebot.Event) {
	_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(creditsInstructions)).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
}
