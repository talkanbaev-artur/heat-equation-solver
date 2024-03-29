package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/config"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/logic"
	"github.com/talkanbaev-artur/heat-equation-solver/src/backend/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	lg := logrus.New()
	config.Init()

	s := logic.NewAPIService()

	r := mux.NewRouter()
	server.MakeMuxRoutes(s, r.PathPrefix("/api").Subrouter(), lg)
	path, _ := os.Getwd()
	path = filepath.Join(path, "frontend")
	server.AttachSPA(r, path, "index.html")
	go func() {
		c := cors.New(cors.Options{AllowedOrigins: []string{"*"}, AllowedMethods: []string{"POST", "GET", "OPTIONS"}})
		srv := http.Server{
			Handler:      c.Handler(r),
			ReadTimeout:  time.Minute,
			WriteTimeout: time.Minute,
			Addr:         fmt.Sprintf("0.0.0.0:%d", config.Config().Port),
		}
		lg.WithField("port", config.Config().Port).Info("The web server is up")
		lg.Error(srv.ListenAndServe().Error())
		cancel()
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		lg.Info("Catched shutdown signal - initiating graceful shutdown")
		cancel()
	}()

	lg.Info("Initialised heat equation solver application")
	<-ctx.Done()
	lg.Info("Shutting down...")
}
