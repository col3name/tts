package http

import (
	"encoding/json"
	"github.com/col3name/tts/pkg/model"
	"github.com/col3name/tts/pkg/repo"
	"io/ioutil"
	"log"
	"net/http"
)

type SettingController struct {
	SettingRepo repo.SettingRepo
}

func (c *SettingController) GetSettings(w http.ResponseWriter, _ *http.Request) {
	settings, err := c.SettingRepo.GetSettings()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	marshal, err := json.Marshal(settings)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(marshal)
	w.WriteHeader(http.StatusOK)
}

func (c *SettingController) SetSettings(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if (*req).Method == "OPTIONS" {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var settingDb model.SettingDB
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(data, &settingDb)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = c.SettingRepo.SaveSettings(&settingDb)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	v := struct {
		Status  int
		Message string
	}{Status: 1, Message: "ok"}
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(marshal)
	w.WriteHeader(http.StatusOK)
}
