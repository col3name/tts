package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/service/voice"
	"sync"
	"time"
)

type ChatListener struct {
	speakService voice.SpeechVoiceService
	ch           chan string
	chCompleted  chan bool
	isFirst      bool
	mx           sync.RWMutex
}

func NewChatListener(service voice.SpeechVoiceService) *ChatListener {
	c := new(ChatListener)
	c.speakService = service
	c.ch = make(chan string, 100)
	c.chCompleted = make(chan bool, 1)
	c.isFirst = true
	return c
}

func (l *ChatListener) Handle() {
	l.isFirst = true
	for {
		select {
		case <-l.chCompleted:
			l.mx.RLock()
			text := <-l.ch
			fmt.Println("start", time.Now(), text)
			err := l.speakService.Speak(text)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("complete", time.Now(), text)
			l.chCompleted <- true
			l.mx.RUnlock()
		}
	}
}

func (l *ChatListener) OnShardMessage(_ int, msg irc.ChatMessage) {
	s := msg.Sender.DisplayName + " say that " + msg.Text + " !"
	fmt.Println("send", s)
	l.ch <- s
	if l.isFirst {
		l.chCompleted <- true
		l.isFirst = false
	}
}
