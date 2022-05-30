package http

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
}

func (s *Server) StartServer(serverUrl string, handler http.Handler) *http.Server {
	log.Info("Started at" + serverUrl)
	srv := &http.Server{Addr: serverUrl, Handler: handler}
	log.Error(srv.ListenAndServe())
	return srv
}

func (s *Server) GetKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)

	return osKillSignalChan
}

func (s *Server) WaitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
