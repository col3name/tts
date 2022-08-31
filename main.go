package main

import (
	"github.com/col3name/tts/pkg/config"
	"github.com/col3name/tts/pkg/handler/twitch"
	"github.com/col3name/tts/pkg/service/application"
	"github.com/col3name/tts/pkg/service/setting"
)

func main() {
	conf := config.NewConfig()
	conf.Parse()

	settingRepo := setting.SetupSettingRepo()
	stng := setting.UpdateSetting(conf, settingRepo)

	messageHandler := twitch.NewMessageHandler(stng, conf, settingRepo)
	app := application.NewApplication(settingRepo, messageHandler)
	go app.RunStaticServer()
	go app.RunWebView()
	go app.RunApiServer(settingRepo, conf.RestAddress)
	app.RunTwitchChatHandler()
}
