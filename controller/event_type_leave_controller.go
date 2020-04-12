package controller

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
)

type EventTypeLeaveController struct {}

func(controller EventTypeLeaveController) Execute(bot *linebot.Client, event *linebot.Event) {
	profile, err := bot.GetProfile(event.Source.UserId).Do()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Left from " + profile.DisplayName + " group")
}
