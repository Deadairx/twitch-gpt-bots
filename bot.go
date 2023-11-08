package main

import (
    "github.com/gempir/go-twitch-irc/v2"
    "github.com/joho/godotenv"
    "os"
    "log"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    client := twitch.NewClient(os.Getenv("BOT_USERNAME"), os.Getenv("TWITCH_OAUTH"))

    client.OnPrivateMessage(func(message twitch.PrivateMessage) {
        if message.Message == "!ping" {
            client.Say(message.Channel, "Pong")
        }

        if message.Message == "!hello" {
            client.Say(message.Channel, "Hello, "+ message.User.DisplayName +"!")
        }
    })

    client.Join(os.Getenv("TWITCH_CHANNEL"))

    err = client.Connect()
    if err != nil {
        panic(err)
    }
}
