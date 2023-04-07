package main

import (
	"os"
	"telegopher"
)

//"fmt"
//"telegopher/core/telegram"

func main() {
	var token string = os.Getenv("BOT_TOKEN")
	//var apiCaller telegopher.ApiCaller = *telegopher.NewApiCaller("https://api.telegram.org/bot/")
	//fmt.Println(apiCaller.GetNonParamsCallApi("getMe"))
	//var bot telegram.Telegram = *telegram.NewTelegram(token)
	// bot.GetUpdates(0, 100, 1000, telegram.AllowedUpdates{"message"})

	//fmt.Printf("%+v\n", bot.GetMe())
	var bot telegopher.Telegopher = *telegopher.NewTelegopher(token)
	bot.AddHandlerMessage(telegopher.NewTrigger("text", "привет"), helloHandler)
	bot.StartLongPolling(&telegopher.BotOpts{})
}

func helloHandler(ctx *telegopher.Context) {
	ctx.Reply("мир", struct{}{})
}
