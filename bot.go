package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    client := twitch.NewClient(os.Getenv("BOT_USERNAME"), os.Getenv("TWITCH_OAUTH"))

    rateLimiter := NewCommandRateLimiter(5 * time.Second)

    // TODO: Read commands from a file

    client.OnPrivateMessage(func(message twitch.PrivateMessage) {
        words := strings.Fields(message.Message)
        if len(words) >= 5 {
            log.Println("OpenAI Trigger: " + message.User.DisplayName + " said " + message.Message)
            if rateLimiter.CanExecute() {
                log.Println("Getting response from OpenAI")
                response, err := GetOpenAIResponse(message.Message)
                if err != nil {
                    log.Println("Error getting response from OpenAI:", err)
                    return
                }

                log.Println("Response from OpenAI:", response)
                client.Say(message.Channel, response)
                rateLimiter.Execute()
            }
        }

        if message.Message == "!ping" {
            log.Println(message.User.DisplayName + " said ping")
            if rateLimiter.CanExecute() {
                log.Println("Pong!")
                client.Say(message.Channel, "Pong")
                rateLimiter.Execute()
            }
        }

        if message.Message == "!hello" {
            log.Println(message.User.DisplayName + " said hello")
            client.Say(message.Channel, "Hello, "+ message.User.DisplayName +"!")
        }

        if message.Message == "!repo" {
            client.Say(message.Channel, "https://github.com/Deadairx/twitch-stream")
        }
    })

    client.Join(os.Getenv("TWITCH_CHANNEL"))

    err = client.Connect()
    if err != nil {
        panic(err)
    }
}

