package main

import (
    "os"
    "strings"
    "log"
    "net/http"
    "encoding/json"
    "bytes"

    tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type resultUrl struct {
	Result_url string
}

func main() {
    bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
    if err != nil {
        panic(err)
    }

    bot.Debug = true
    // Create a new UpdateConfig struct with an offset of 0. Offsets are used
    // to make sure Telegram knows we've handled previous values and we don't
    // need them repeated.
    updateConfig := tgbotapi.NewUpdate(0)

    // Tell Telegram we should wait up to 30 seconds on each request for an
    // update. This way we can get information just as quickly as making many
    // frequent requests without having to send nearly as many.
    updateConfig.Timeout = 30

    // Start polling Telegram for updates.
    updates := bot.GetUpdatesChan(updateConfig)
    
    var count int
    urls := make([]string, 3)

    // Let's go through each update that we're getting from Telegram.
    for update := range updates {
        // Telegram can send many types of updates depending on what your Bot
        // is up to. We only want to look at messages for now, so we can
        // discard any other updates.
        if update.Message == nil {
            continue
        }

        if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		cfg := tgbotapi.NewSetMyCommands(
			tgbotapi.BotCommand{
				Command:     "/help",
				Description: "Help command",
			},
			tgbotapi.BotCommand{
				Command: "/dummy",
				Description: "stub command",
			},
		)
		_, err := bot.Request(cfg)
		if err != nil {
			log.Fatal(err)	
		}

		// TODO: rewrite to use proper commands as command link to article 
        	// Extract the command from the Message.
        	switch update.Message.Command() {
        	case "help":
            		msg.Text = "Bot understand as text 'ADD <url>'."
        	default:
            		msg.Text = "I don't know that command"
        	}

        	if _, err := bot.Send(msg); err != nil {
            		log.Panic(err)
        	}
		continue
        }

	commands := strings.Split(update.Message.Text, " ")

	// TODO: could be a better way to create NewPoll from commands 
	switch commands[0] {
	case "ADD":
		if len(commands) == 2 {
			if count == 2 {
				urls[2] = urlShort("url=" + commands[1])
				// NewPoll(chatID int64 with dash)
				bot.Send(tgbotapi.NewPoll(-604350070, "голосуем за лучшую статью:", urls[0], urls[1], urls[2]))
				count = 0
				break
			}
			if count == 0 {
				urls[0] = urlShort("url=" + commands[1])
				count++
				break
			}
			if count == 1 {
				urls[1] = urlShort("url=" + commands[1])
				count++
				break
			}
		}
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Wrong Command"))
	default:
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Command not Found"))
	}
    }
}

// cleanuri has request limit so it might be not 100% reliable
func urlShort(url string) string {
	var urlStr = []byte(url)
	response, err := http.Post("https://cleanuri.com/api/v1/shorten", "application/x-www-form-urlencoded",
        bytes.NewBuffer(urlStr))

    	if err != nil {
        	log.Fatal(err)
    	}

    	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var res resultUrl
	err = decoder.Decode(&res)

	if err != nil {
		panic(err)
	}

	return res.Result_url
}

