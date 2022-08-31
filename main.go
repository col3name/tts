package main

import (
	"context"
	"database/sql"
	"github.com/Adeithe/go-twitch"
	"github.com/Adeithe/go-twitch/irc"
	"github.com/col3name/tts/pkg/config"
	"github.com/col3name/tts/pkg/handler"
	"github.com/col3name/tts/pkg/http/transport"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	"github.com/col3name/tts/pkg/repo/sqlite"
	langdetection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/service/voice"
	"github.com/inkeliz/gowebview"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.NewConfig()
	conf.Parse()

	settingRepo := setupSettingRepo()
	setting := updateSetting(conf, settingRepo)
	messageHandler := newTwitchMessageHandler(setting, conf, settingRepo)
	app := newApplication(settingRepo, messageHandler)
	go app.runStaticServer()
	go app.runWebView()
	go app.runApiServer(settingRepo, conf.RestAddress)
	app.runTwitchChatHandler()
}

func setupChatListener(setting *model.SettingDB, config *config.Config, settingRepo repo.SettingRepo) handler.ChatListener {
	messageFilter := moderation.NewMessageFilter(setting.ReplacementWordPair, setting.IgnoreWords, config.UserBanList)
	linguaDetectionService := langdetection.NewLinguaDetectionService(langdetection.DefaultLanguages)
	ttsService := voice.NewGoTtsService(setting.Language, messageFilter, setting.Volume, settingRepo, setting.LanguageDetectorEnabled, linguaDetectionService)
	return handler.NewChatListener(ttsService)
}

type application struct {
	settingRepo          repo.SettingRepo
	twitchMessageHandler *twitchMessageHandler
	conf                 *config.Config
}

func newApplication(settingRepo repo.SettingRepo, messageHandler *twitchMessageHandler) *application {
	return &application{settingRepo: settingRepo, twitchMessageHandler: messageHandler}
}
func (a *application) runStaticServer() {
	fs := http.FileServer(http.Dir("./web/build"))
	http.Handle("/", fs)

	log.Print("Listening on ", a.conf.StaticApiAddress)
	err := http.ListenAndServe(a.conf.StaticApiAddress, nil)
	logFatalIfNeed(err)
}

func (a *application) runWebView() {
	webviewConfig := gowebview.WindowConfig{Title: "Text to speech",
		Size: &gowebview.Point{X: 800, Y: 800}}

	webView, err := gowebview.New(&gowebview.Config{URL: a.conf.WebViewAddress, WindowConfig: &webviewConfig})
	logFatalIfNeed(err)

	defer webView.Destroy()
	webView.Run()
}

func (a *application) runApiServer(settingRepo repo.SettingRepo, serveAddress string) {
	router := transport.NewRouter(settingRepo)
	server := transport.Server{}
	killSignalChan := server.GetKillSignalChan()
	httpServer := server.StartServer(serveAddress, router)

	server.WaitForKillSignal(killSignalChan)
	err := httpServer.Shutdown(context.Background())
	logFatalIfNeed(err)
}

type twitchMessageHandler struct {
	setting     *model.SettingDB
	config      *config.Config
	settingRepo repo.SettingRepo
	shards      *irc.Client
}

func newTwitchMessageHandler(setting *model.SettingDB, config *config.Config, settingRepo repo.SettingRepo) *twitchMessageHandler {
	return &twitchMessageHandler{
		setting:     setting,
		config:      config,
		settingRepo: settingRepo,
		shards:      twitch.IRC(),
	}
}

func (t *twitchMessageHandler) Handle() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)
	t.run()
	<-osSignals
	t.stop()
}

func (t *twitchMessageHandler) run() {
	chatListener := setupChatListener(t.setting, t.config, t.settingRepo)
	t.shards.OnShardMessage(chatListener.OnShardMessage)
	go chatListener.Handle()

	log.Println("Started")
	err := t.shards.Join(t.config.ChannelsList...)
	logFatalIfNeed(err)
}

func (t *twitchMessageHandler) stop() {
	err := os.Remove(voice.AudioFolder)
	if err != nil {
		log.Println(err)
	}
	log.Println("Stopping...")
	t.shards.Close()
}

func (a *application) runTwitchChatHandler() {
	a.twitchMessageHandler.Handle()
}

func updateSetting(config *config.Config, settingRepo repo.SettingRepo) *model.SettingDB {
	setting := setupSettingFromConfig(config)

	settingsFromDatabase, err := settingRepo.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	mergeSetting(setting, settingsFromDatabase)
	if err = settingRepo.SaveSettings(setting); err != nil {
		log.Fatal(err)
	}
	return setting
}

func mergeSetting(setting *model.SettingDB, settingsFromDatabase *model.SettingDB) {
	if settingsFromDatabase != nil {
		setting.ReplacementWordPair = settingsFromDatabase.ReplacementWordPair
		setting.IgnoreWords = settingsFromDatabase.IgnoreWords
		setting.UserBanList = settingsFromDatabase.UserBanList
		setting.ChannelsToListen = settingsFromDatabase.ChannelsToListen
	}
}

func setupSettingRepo() repo.SettingRepo {
	db, err := sql.Open("sqlite3", "./data.db")
	logFatalIfNeed(err)
	settingRepo, err := sqlite.NewSettingRepoImpl(db)
	logFatalIfNeed(err)
	return settingRepo
}

func setupSettingFromConfig(config *config.Config) *model.SettingDB {
	s := &model.SettingDB{
		Id:                      1,
		ReplacementWordPair:     config.ModerationWordPairs,
		IgnoreWords:             config.IgnoreWords,
		Language:                config.Language,
		LanguageDetectorEnabled: config.LangDetectorEnabled,
		ChannelsToListen:        "",
		Volume:                  1,
	}
	s.SetUserBanList(config.UserBanList)
	return s
}

func logFatalIfNeed(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
