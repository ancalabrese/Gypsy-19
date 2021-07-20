package main

import (
	// "fmt"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/ancalabrese/Gypsy-19/Scraper/Data"
	"github.com/ancalabrese/Gypsy-19/Scraper/Handler"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
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

	go func() {
		l.Info("Starting server instance", "Name", conf.Scraper.ServerName, "URL", conf.Scraper.Url)
		s := Handler.CreateScraper(conf, l)
	retry:
		err := s.RetrieveLists()
		if err != nil {
			goto retry
		}
		ctx := context.Background()
		ghToken, _ := viper.Get("GITHUB_TOKEN").(string)
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
		httpClient := oauth2.NewClient(ctx, ts)
		ghc := Data.NewGitHubConnector(l, httpClient, &conf.Connector.CommitterInfo, &conf.Connector.Repo)
		ghc.UpdateDB(s.TravelLists)
	}()

	signChannel := make(chan os.Signal)
	signal.Notify(signChannel, os.Interrupt)
	signal.Notify(signChannel, os.Kill)

	sig := <-signChannel
	l.Info("Received system signal", "sig", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ctx.Done()
}
