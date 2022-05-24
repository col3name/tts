package handler

import (
	"fmt"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/service"
	"time"
)

type ChatListener struct {
	speakService service.SpeakService
}

func NewChatListener(service service.SpeakService) *ChatListener {
	c := new(ChatListener)
	c.speakService = service
	return c
}

func (l *ChatListener) OnShardReconnect(shardID int) {
	fmt.Printf("Shard #%d reconnected\n", shardID)
}

func (l *ChatListener) OnShardLatencyUpdate(shardID int, latency time.Duration) {
	fmt.Printf("Shard #%d has %dms ping\n", shardID, latency.Milliseconds())
}

func (l *ChatListener) OnShardMessage(shardID int, msg irc.ChatMessage) {
	fmt.Println(msg)

	text := msg.Sender.DisplayName + "говорит что" + msg.Text + " placeholder"
	fmt.Println(text)
	err := l.speakService.Speak(text)
	if err != nil {
		fmt.Println(err)
	}
}
