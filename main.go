package main

import (
	"context"
	"fmt"
	"github.com/Adeithe/go-twitch"
	"github.com/col3name/tts/pkg/handler"
	httpTranposrt "github.com/col3name/tts/pkg/http"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	lang_detection "github.com/col3name/tts/pkg/service/lang-detection"
	"github.com/col3name/tts/pkg/service/moderation"
	"github.com/col3name/tts/pkg/service/voice"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	channels := os.Getenv("TWITCH_CHANNEL")
	if len(channels) == 0 {
		log.Fatal("Error loading .env file")
	}
	channelsList := strings.Split(channels, ",")
	language := os.Getenv("LANGUAGE")
	if !voice.IsSupported(language) {
		log.Fatal("Not supported language. ")
	}
	moderationPair := os.Getenv("MODERATION")
	ignoreString := os.Getenv("IGNORE")
	volumeString := os.Getenv("VOLUME")
	userIgnore := os.Getenv("USER_IGNORE")
	langDetectorEnabledString := os.Getenv("LANG_DETECT_ENABLED")
	langDetectorEnabled := false
	if langDetectorEnabledString == "" || langDetectorEnabledString == "false" {
		langDetectorEnabled = false
	} else if langDetectorEnabledString == "true" {
		langDetectorEnabled = true
	}
	usersList := strings.Split(userIgnore, ",")
	usersList = append(usersList, channelsList...)
	volume := 7
	if len(volumeString) != 0 {
		num, err := strconv.Atoi(volumeString)
		ifNeedFatal(err)
		volume = num
	}
	settingDB := model.SettingDB{
		Id:                      1,
		ReplacementWordPair:     moderationPair,
		IgnoreWords:             ignoreString,
		Language:                language,
		LanguageDetectorEnabled: langDetectorEnabled,
		UserBanList: func(list []string) string {
			res := ""
			for _, s := range list {
				res += s
			}
			return res
		}(usersList),
		ChannelsToListen: "",
		Volume:           1,
	}
	config, err := ParseConfig()
	ifNeedFatal(err)
	fmt.Println(config.ServeRestAddress)

	connector, err := repo.GetConnector(config.Config)
	ifNeedFatal(err)

	pool, err := repo.NewConnectionPool(connector)
	ifNeedFatal(err)
	settingRepo := repo.NewSettingRepo(pool)
	db, err := settingRepo.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	settingDB.ReplacementWordPair = db.ReplacementWordPair
	settingDB.IgnoreWords = db.IgnoreWords
	settingDB.UserBanList = db.UserBanList
	settingDB.ChannelsToListen = db.ChannelsToListen

	if err = settingRepo.SaveSettings(&settingDB); err != nil {
		log.Fatal(err)
	}
	go func(setting *model.SettingDB, settingRepo repo.SettingRepo) {
		router := httpTranposrt.NewRouter(settingRepo)
		server := httpTranposrt.Server{}
		killSignalChan := server.GetKillSignalChan()
		srv := server.StartServer(config.ServeRestAddress, router)

		server.WaitForKillSignal(killSignalChan)
		err = srv.Shutdown(context.Background())
		ifNeedFatal(err)
	}(&settingDB, settingRepo)

	filter := moderation.NewDefaultFilter(moderationPair, ignoreString, usersList)
	detectionService := lang_detection.NewLinguaDetectionService(lang_detection.DefaultLanguages)
	service := voice.NewHtgoTtsService(language, filter, volume, settingRepo, langDetectorEnabled, detectionService)
	chatListener := handler.NewChatListener(service)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	shards := twitch.IRC()
	shards.OnShardMessage(chatListener.OnShardMessage)
	go chatListener.Handle()
	log.Println("Started")
	if err := shards.Join(channelsList...); err != nil {
		log.Fatal(err)
	}
	<-sc
	err = os.Remove("audio")
	fmt.Println(err)
	fmt.Println("Stopping...")
	shards.Close()
}

func ifNeedFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
