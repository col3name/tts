package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/service/voice"
	"sync"
	"time"
)

type ChatListener interface {
	Handle()
	OnShardMessage(_ int, message irc.ChatMessage)
}

type TwitchChatListener struct {
	speakService    voice.SpeechVoiceService
	messagesChannel chan model.Message
	chCompleted     chan bool
	isFirst         bool
	mx              sync.RWMutex
}

func NewChatListener(service voice.SpeechVoiceService) *TwitchChatListener {
	c := new(TwitchChatListener)
	c.speakService = service
	c.messagesChannel = make(chan model.Message, 100)
	c.chCompleted = make(chan bool, 1)
	c.isFirst = true
	return c
}

func (l *TwitchChatListener) Handle() {
	l.isFirst = true
	for range l.chCompleted {
		l.mx.RLock()
		message := <-l.messagesChannel
		fmt.Println("start", time.Now(), message)
		err := l.speakService.Speak(message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("complete", time.Now(), message)
		l.chCompleted <- true
		l.mx.RUnlock()
	}
}

func (l *TwitchChatListener) OnShardMessage(_ int, message irc.ChatMessage) {
	text := message.Sender.DisplayName + " say that " + message.Text + " !"
	fmt.Println("send", text)
	l.messagesChannel <- model.Message{From: message.Sender.Username, Text: text}
	if l.isFirst {
		l.chCompleted <- true
		l.isFirst = false
	}
}
