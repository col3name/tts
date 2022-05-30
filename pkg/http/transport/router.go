package transport

import (
	"fmt"
	"github.com/col3name/tts/pkg/repo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NewRouter(settingRepo repo.SettingRepo) http.Handler {
	router := mux.NewRouter()
	controller := SettingController{
		SettingRepo: settingRepo,
	}
	api := router.PathPrefix("/api/v1").Subrouter()
	router.HandleFunc("/health", healthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/ready", readyCheckHandler).Methods(http.MethodGet)
	api.HandleFunc("/settings", controller.GetSettings).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/settings", controller.SetSettings).Methods(http.MethodPost, http.MethodOptions)
	return logMiddleware(router)
}

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "{\"status\": \"OK\"}")
}

func readyCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "{\"host\": \"%v\"}", r.Host)
}

func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":     r.Method,
			"url":        r.URL,
			"remoteAddr": r.RemoteAddr,
			"userAgent":  r.UserAgent(),
		}).Info("got a new request")
		h.ServeHTTP(w, r)
	})
}
