package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Adeithe/go-twitch"
	"github.com/col3name/tts/pkg/handler"
	transport2 "github.com/col3name/tts/pkg/http/transport"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	"github.com/col3name/tts/pkg/repo/sqlite"
	langdetection "github.com/col3name/tts/pkg/service/lang-detection"
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
	serveRestAddress := os.Getenv("SERVE_REST_ADDRESS")
	if len(serveRestAddress) == 0 {
		serveRestAddress = ":8000"
	}

	db, err := sql.Open("sqlite3", "./data.db")
	ifNeedFatal(err)
	settingRepo, err := sqlite.NewSettingRepoImpl(db)
	ifNeedFatal(err)
	out, err := settingRepo.GetSettings()
	if err != nil {
		log.Fatal(err)
	}
	if out != nil {
		settingDB.ReplacementWordPair = out.ReplacementWordPair
		settingDB.IgnoreWords = out.IgnoreWords
		settingDB.UserBanList = out.UserBanList
		settingDB.ChannelsToListen = out.ChannelsToListen
	}
	if err = settingRepo.SaveSettings(&settingDB); err != nil {
		log.Fatal(err)
	}
	go func(settingRepo repo.SettingRepo, serveRestAddress string) {
		router := transport2.NewRouter(settingRepo)
		server := transport2.Server{}
		killSignalChan := server.GetKillSignalChan()
		srv := server.StartServer(serveRestAddress, router)

		server.WaitForKillSignal(killSignalChan)
		err = srv.Shutdown(context.Background())
		ifNeedFatal(err)
	}(settingRepo, serveRestAddress)

	filter := moderation.NewDefaultFilter(moderationPair, ignoreString, usersList)
	detectionService := langdetection.NewLinguaDetectionService(langdetection.DefaultLanguages)
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
