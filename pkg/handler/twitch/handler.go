package twitch

import (
	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/config"
	"github.com/col3name/tts/pkg/handler"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	langdetection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/service/voice"
	"github.com/col3name/tts/pkg/util/logger"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type MessageHandler struct {
	setting     *model.SettingDB
	config      *config.Config
	settingRepo repo.SettingRepo
	shards      *irc.Client
}

func NewMessageHandler(setting *model.SettingDB, config *config.Config, settingRepo repo.SettingRepo) *MessageHandler {
	return &MessageHandler{
		setting:     setting,
		config:      config,
		settingRepo: settingRepo,
		shards:      twitch.IRC(),
	}
}

func (t *MessageHandler) Handle() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	t.run()
	<-osSignals
	t.stop()
}

func (t *MessageHandler) run() {
	chatListener := t.setupChatListener()
	t.shards.OnShardMessage(chatListener.OnShardMessage)
	go chatListener.Handle()

	log.Println("Started")
	err := t.shards.Join(t.config.ChannelsList...)
	logger.LogFatalIfNeed(err)
}

func (t *MessageHandler) stop() {
	err := os.Remove(voice.AudioFolder)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stopping...")
	t.shards.Close()
}

func (t *MessageHandler) setupChatListener() handler.ChatListener {
	settingRepo := t.settingRepo
	setting := t.setting
	banList := t.config.UserBanList

	messageFilter := moderation.NewMessageFilter(setting.ReplacementWordPair, setting.IgnoreWords, banList)
	linguaDetectionService := langdetection.NewLinguaDetectionService(langdetection.DefaultLanguages)
	ttsService := voice.NewGoTtsService(setting.Language, messageFilter, setting.Volume, settingRepo, setting.LanguageDetectorEnabled, linguaDetectionService)
	return handler.NewChatListener(ttsService, t.config.GreetingText)
}
