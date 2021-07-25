package WebScraper

import (
	"github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
)

type Scraper struct {
	url         string
	selector    string
	TravelLists Country.Lists
	logger      hclog.Logger
}

type IScraper interface {
	CreateScraper(config *Settings.Configurations, l hclog.Logger) *Scraper
	RetrieveLists() error
}
