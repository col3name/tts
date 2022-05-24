package main

import (
	"fmt"
	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/handler"
	"github.com/col3name/tts/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token := os.Getenv("TWITCH_TOKEN")
	if len(token) == 0 {
		token = "06km8hcuiad6abryvdsi01b7rm1ghs"
	}
	channels := os.Getenv("TWITCH_CHANNEL")
	if len(token) == 0 {
		channels = "Spiinlock"
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	writer := &irc.Conn{}
	err = writer.SetLogin(channels, token)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := writer.Connect(); err != nil {
		panic("failed to start writer")
	}

	chatListener := handler.NewChatListener(service.NewHtgoTtsService())

	reader := twitch.IRC()
	reader.OnShardReconnect(chatListener.OnShardReconnect)
	reader.OnShardLatencyUpdate(chatListener.OnShardLatencyUpdate)
	reader.OnShardMessage(chatListener.OnShardMessage)

	if err := reader.Join(channels); err != nil {
		panic(err)
	}
	fmt.Println("Connected to IRC!")

	<-sc
	err = os.Remove("audio")
	fmt.Println(err)
	fmt.Println("Stopping...")
	reader.Close()
	writer.Close()
}
