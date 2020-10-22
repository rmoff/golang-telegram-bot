package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	var resp string
	var chatID int64

	// Authorise and create bot instance
	bot, err := tgbotapi.NewBotAPI(TELEGRAM_API_TOKEN)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s (https://t.me/%s)", bot.Self.UserName, bot.Self.UserName)

	// Subscribe to updates
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)

	// Process any messages that we're sent as they arrive
	// on the updates channel
	for update := range updates {
		if update.Message == nil {
			continue
		}

		chatID = update.Message.Chat.ID
		t := update.Message.Text
		log.Printf("[%s] %s (command: %v, location: %v)\n", update.Message.From.UserName, t, update.Message.IsCommand(), update.Message.Location)
		switch {
		case update.Message.IsCommand():
			// Handle commands
			switch update.Message.Command() {
			case "start", "help":
				bot.Send(tgbotapi.NewMessage(chatID, "🤖 Here I am, brain the size of a planet, and they make me an example bot in a conference talk. Call that job satisfaction? 'Cos I don't.\n\nOh yeah, and I've started the bot now."))
				// do stuff for the "alert" command
			case "alert":
				bot.Send(tgbotapi.NewMessage(chatID, "🚨 alert! alert! are you a lert? what even is a lert? 🤔"))
				// catch-all command handling
			default:
				bot.Send(tgbotapi.NewMessage(chatID, "🤔 Command not recognised."))
			}
		case update.Message.Location != nil:
			// They sent us a location
			resp = "Ooh nice location, I think my cat lives near there 😸"
			msg := tgbotapi.NewMessage(chatID, resp)
			if _, e := bot.Send(msg); e != nil {
				log.Printf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, e)
			}
		default:
			resp = fmt.Sprintf("\"%v\" Right back atcha! 💥", t)
			msg := tgbotapi.NewMessage(chatID, resp)

			if _, e := bot.Send(msg); e != nil {
				log.Printf("Error sending message to telegram.\nMessage: %v\nError: %v", msg, e)
			}
		}

	}
}
