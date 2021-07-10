package main

import (
	// "fmt"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/Gypsy-19/Server/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
)

func main() {

	l := hclog.New(&hclog.LoggerOptions{
		Level: hclog.Debug,
	})
	//Load service config
	conf, err :=  Settings.Load("config.yml", l)

	if err != nil {
		l.Error("Cannot start server quitting.")
		os.Exit(1)
	}

	go func() {
		l.Info("Starting server instance", "Name", conf.Scraper.ServerName, "URL", conf.Scraper.Url)
	}()

	signChannel := make(chan os.Signal)
	signal.Notify(signChannel, os.Interrupt)
	signal.Notify(signChannel, os.Kill)

	sig := <-signChannel
	l.Info("Received system signal", "sig", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ctx.Done()
}
