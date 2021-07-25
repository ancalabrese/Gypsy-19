package WebScraper

import (
	"errors"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
)

func CreateScraper(config *Settings.Configurations, l hclog.Logger) *Scraper {
	s := &Scraper{
		url:         config.Scraper.Url,
		selector:    config.Scraper.HtmlListSelector,
		TravelLists: Country.Lists{},
		logger:      l,
	}
	s.logger.Debug("Created new scraper", "url", s.url)
	return s
}

func (s *Scraper) RetrieveLists() error {
	s.logger.Info("Starting scraping remote resource")
	err := s.getPage()
	if err != nil {
		s.logger.Error("Failed retrieving list", "error", err)
		return err
	}
	s.logger.Info("Finished scraping remote resource")
	return nil
}

func (p *Scraper) getPage() error {
	p.logger.Debug("Visiting page", "url", p.url)
	httpResp, err := http.Get(p.url)
	if err != nil {
		p.logger.Error("Scraping failed", "Error", err)
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		p.logger.Error("Cannot visit "+p.url, "Status code", httpResp.StatusCode, "Status", httpResp.Status)
		return errors.New("Canot access url. Status code " + httpResp.Status + " Status: " + httpResp.Status)
	}
	p.extractCuntryLists(httpResp.Body)
	p.logger.Debug("Retrieved countries", "#", len(p.TravelLists.Amber)+len(p.TravelLists.Green)+len(p.TravelLists.Red))
	return nil
}

func (p *Scraper) extractCuntryLists(body io.Reader) error {
	p.logger.Debug("Extracting countries")
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		p.logger.Error("Error loading HTTP page", "Error", err)
		return err
	}

	var countries []Country.Country
	doc.Find(p.selector).Each(func(i int, s *goquery.Selection) {
		countries = append(countries, Country.Country{
			Name: s.Text(),
		})
	})

	// switch list {
	// case "red":
	// 	p.TravelLists.Red = countries
	// case "amber":
	// 	p.TravelLists.Amber = countries
	// case "green":
	// 	p.TravelLists.Green = countries
	// }
	return nil
}