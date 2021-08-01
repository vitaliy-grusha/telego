package main

import (
	"fmt"
	"os"

	telego "github.com/mymmrac/go-telegram-bot-api"
)

func main() {
	botToken := os.Getenv("TOKEN")

	bot, err := telego.NewBot(botToken)
	if err != nil {
		fmt.Println(err)
		return
	}

	bot.DebugMode(true)

	// Setup a webhook
	_ = bot.SetWebhook(&telego.SetWebhookParams{
		URL:         "https://www.google.com:443/" + bot.Token(),
		Certificate: &telego.InputFile{File: mustOpen("cert.pem")},
	})

	// Receive information about webhook
	info, _ := bot.GetWebhookInfo()
	fmt.Println()
	fmt.Printf("%#v\n", info)
	fmt.Println()

	// Get updates channel from webhook
	updates, _ := bot.ListenForWebhook("/" + bot.Token())

	// Start server for receiving requests from telegram
	bot.StartListeningForWebhook("0.0.0.0:443/"+bot.Token(), "cert.pem", "key.pem")

	// Loop through all updates when they came
	for update := range updates {
		fmt.Println("====")
		fmt.Printf("%#v\n", update)
		fmt.Println("====")
	}
}

// Helper function to open file or panic
func mustOpen(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
