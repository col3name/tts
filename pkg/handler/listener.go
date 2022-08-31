package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/service/voice"
	"sync"
	"time"
)

type ChatListener struct {
	speakService voice.SpeechVoiceService
	ch           chan model.Message
	chCompleted  chan bool
	isFirst      bool
	mx           sync.RWMutex
}

func NewChatListener(service voice.SpeechVoiceService) *ChatListener {
	c := new(ChatListener)
	c.speakService = service
	c.ch = make(chan model.Message, 100)
	c.chCompleted = make(chan bool, 1)
	c.isFirst = true
	return c
}

func (l *ChatListener) Handle() {
	l.isFirst = true
	for range l.chCompleted {
		l.mx.RLock()
		message := <-l.ch
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

func (l *ChatListener) OnShardMessage(_ int, message irc.ChatMessage) {
	s := message.Sender.DisplayName + " say that " + message.Text + " !"
	fmt.Println("send", s)
	l.ch <- model.Message{From: message.Sender.Username, Text: s}
	if l.isFirst {
		l.chCompleted <- true
		l.isFirst = false
	}
}
