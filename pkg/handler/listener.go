package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/service/voice"
)

type ChatListener struct {
	speakService voice.SpeechVoiceService
}

func NewChatListener(service voice.SpeechVoiceService) *ChatListener {
	c := new(ChatListener)
	c.speakService = service
	return c
}

func (l *ChatListener) OnShardMessage(_ int, msg irc.ChatMessage) {
	text := msg.Sender.DisplayName + " говорит что " + msg.Text + " placeholder"
	err := l.speakService.Speak(text)
	if err != nil {
		fmt.Println(err)
	}
}
