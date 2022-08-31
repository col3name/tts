package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/service/voice"
	"github.com/col3name/tts/pkg/util/separator"
	"sync"
	"time"
)

type ChatListener interface {
	Handle()
	OnShardMessage(_ int, message irc.ChatMessage)
}

type TwitchChatListener struct {
	isFirst         bool
	greetingText    string
	mx              sync.RWMutex
	speakService    voice.SpeechVoiceService
	messagesChannel chan model.Message
	chCompleted     chan bool
}

func NewChatListener(service voice.SpeechVoiceService, greetingText string) *TwitchChatListener {
	c := new(TwitchChatListener)
	c.speakService = service
	c.messagesChannel = make(chan model.Message, 100)
	c.chCompleted = make(chan bool, 1)
	c.isFirst = true
	c.greetingText = greetingText
	return c
}

func (t *TwitchChatListener) Handle() {
	t.isFirst = true
	for range t.chCompleted {
		t.mx.RLock()
		message := <-t.messagesChannel
		fmt.Println("start", time.Now(), message)
		err := t.speakService.Speak(message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("complete", time.Now(), message)
		t.chCompleted <- true
		t.mx.RUnlock()
	}
}

func (t *TwitchChatListener) OnShardMessage(_ int, message irc.ChatMessage) {
	text := message.Sender.DisplayName + separator.Space + t.greetingText + separator.Space + message.Text + " !"
	fmt.Println("send", text)
	t.messagesChannel <- model.Message{From: message.Sender.Username, Text: text}
	if t.isFirst {
		t.chCompleted <- true
		t.isFirst = false
	}
}
