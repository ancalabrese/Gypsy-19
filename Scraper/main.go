package main

import (
	// "fmt"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/Gypsy-19/Scraper/DB/GitHubClient"
	"github.com/ancalabrese/Gypsy-19/Scraper/Handler"
	"github.com/ancalabrese/Gypsy-19/Scraper/Middleware"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/ancalabrese/Gypsy-19/Scraper/WebScraper"
	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

func main() {

	l := hclog.New(&hclog.LoggerOptions{
		Level: hclog.Debug,
	})
	//Load service config
	conf, err := Settings.Load("config.yml", l)
	if err != nil {
		l.Error("Cannot load Server Config", "Error", err)
		os.Exit(1)
	}
	viper.SetConfigFile("local.env")
	err = viper.ReadInConfig()
	if err != nil {
		l.Error("Failed to read config file", "Error", err)
		os.Exit(1)
	}

	ghConnector := GitHubClient.NewGitHubConnector(l, &conf.Connector.CommitterInfo, &conf.Connector.Repo)
	webScraper := WebScraper.CreateScraper(conf, l)
	tlUpdateHandler := Handler.NewTravelListUpdateHandler(l, ghConnector, conf, *webScraper)
	//main API router
	r := mux.NewRouter()
	middlewareLogger := middleware.NewLogger(l)
	r.Use(middlewareLogger.LogIncomingReq)

	//Travel List router
	tlRouter := r.NewRoute().PathPrefix(conf.Service.ApiBasePath + "/lists/update").Subrouter()
	getTlRouter := tlRouter.Methods(http.MethodGet).Subrouter()
	getTlRouter.HandleFunc("", tlUpdateHandler.UpdateTravelInfo)

	//CORS
	corsHandler := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{conf.Service.CorsAllowedOrigins}))

	s := &http.Server{
		Addr:         conf.Service.Url + ":" + conf.Service.Port,
		Handler:      corsHandler(r),
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  conf.Service.IdleTimeout * time.Second,
		ReadTimeout:  conf.Service.ReadTimeout * time.Second,
		WriteTimeout: conf.Service.WriteTimeout * time.Second,
	}

	go func() {
		l.Info("Starting server", "Address", s.Addr)
		err := s.ListenAndServe()
		if err != nil {
			l.Error("Error starting server", "error", err)
			os.Exit(1)
		}
	}()

	signChannel := make(chan os.Signal)
	signal.Notify(signChannel, os.Interrupt)
	signal.Notify(signChannel, os.Kill)

	sig := <-signChannel
	l.Info("Received system signal", "sig", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ctx.Done()
}
