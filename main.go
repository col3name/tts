package main

import (
	"fmt"
	"github.com/Adeithe/go-twitch"
	"github.com/col3name/tts/pkg/handler"
	"github.com/col3name/tts/pkg/service/moderation"
	"strings"

	//"github.com/col3name/tts/pkg/handler"
	//"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/service/voice"
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

	channels := os.Getenv("TWITCH_CHANNEL")
	if len(channels) == 0 {
		log.Fatal("Error loading .env file")
	}
	channelsList := strings.Split(channels, ",")

	fmt.Println(channels)
	language := os.Getenv("LANGUAGE")
	if !voice.IsSupported(language) {
		log.Fatal("Not supported language. ")
	}
	moderationPair := os.Getenv("MODERATION")
	ignoreString := os.Getenv("IGNORE")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	wordsFilter := moderation.NewFilterDefault(moderationPair, ignoreString)

	chatListener := handler.NewChatListener(voice.NewHtgoTtsService(language, wordsFilter))
	shards := twitch.IRC()
	shards.OnShardMessage(chatListener.OnShardMessage)
	log.Println("Started")
	if err := shards.Join(channelsList...); err != nil {
		panic(err)
	}

	<-sc
	err = os.Remove("audio")
	fmt.Println(err)
	fmt.Println("Stopping...")
	shards.Close()
}
