package application

import (
	"context"
	"github.com/col3name/tts/pkg/config"
	"github.com/col3name/tts/pkg/handler/twitch"
	"github.com/col3name/tts/pkg/http/transport"
	"github.com/col3name/tts/pkg/repo"
	"github.com/col3name/tts/pkg/util/logger"
	"github.com/inkeliz/gowebview"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Application struct {
	settingRepo          repo.SettingRepo
	twitchMessageHandler *twitch.MessageHandler
	conf                 *config.Config
}

func NewApplication(settingRepo repo.SettingRepo, messageHandler *twitch.MessageHandler) *Application {
	return &Application{settingRepo: settingRepo, twitchMessageHandler: messageHandler}
}

func (a *Application) RunStaticServer() {
	fs := http.FileServer(http.Dir("./web/build"))
	http.Handle("/", fs)

	log.Print("Listening on ", a.conf.StaticApiAddress)
	err := http.ListenAndServe(a.conf.StaticApiAddress, nil)
	logger.LogFatalIfNeed(err)
}

func (a *Application) RunWebView() {
	webviewConfig := gowebview.WindowConfig{Title: "Text to speech",
		Size: &gowebview.Point{X: 800, Y: 800}}

	webView, err := gowebview.New(&gowebview.Config{URL: a.conf.WebViewAddress, WindowConfig: &webviewConfig})
	logger.LogFatalIfNeed(err)

	defer webView.Destroy()
	webView.Run()
}

func (a *Application) RunApiServer(settingRepo repo.SettingRepo, serveAddress string) {
	router := transport.NewRouter(settingRepo)
	server := transport.Server{}
	killSignalChan := server.GetKillSignalChan()
	httpServer := server.StartServer(serveAddress, router)

	server.WaitForKillSignal(killSignalChan)
	err := httpServer.Shutdown(context.Background())
	logger.LogFatalIfNeed(err)
}

func (a *Application) RunTwitchChatHandler() {
	a.twitchMessageHandler.Handle()
}
