package Handler

import (
	"net/http"

	"github.com/ancalabrese/Gypsy-19/Scraper/DB"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/ancalabrese/Gypsy-19/Scraper/WebScraper"
	"github.com/hashicorp/go-hclog"
)

type TravelListUpdate struct {
	l hclog.Logger
	c DB.DbCLient
	config *Settings.Configurations
	webScraper WebScraper.Scraper
}

func NewTravelListUpdateHandler(logger hclog.Logger, dbc DB.DbCLient, conf *Settings.Configurations, webScraper WebScraper.Scraper) *TravelListUpdate {
	return &TravelListUpdate{l: logger, c: dbc, config: conf, webScraper: webScraper}
}

func (tlu *TravelListUpdate) UpdateTravelInfo(rw http.ResponseWriter, r *http.Request) {
	tlu.l.Info("Got new update list request")
	err :=  tlu.webScraper.RetrieveLists()
	
	if err != nil {
		tlu.l.Error("Returning error", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tlu.c.UpdateDB(tlu.webScraper.TravelLists)
	if err != nil {
		tlu.l.Error("Returning error", "error", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
