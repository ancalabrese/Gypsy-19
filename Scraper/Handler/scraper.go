package Handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
)

type Scraper struct {
	Url         string
	Selectors   map[string]string
	TravelLists Country.Lists
	logger      hclog.Logger
}

func CreateScraper(config *Settings.Configurations, l hclog.Logger) *Scraper {
	s := &Scraper{
		Url:         config.Scraper.Url,
		Selectors:   config.Scraper.HtmlListSelectors,
		TravelLists: Country.Lists{},
		logger:      l,
	}
	s.logger.Debug("Created new scraper", "url", s.Url)
	return s
}

func (p *Scraper) RetrieveLists() error {
	p.logger.Info("Starting scraping remote resource")
	err := p.getPage()
	if err != nil {
		p.logger.Error("Failed retrieving list", "error", err)
		return err
	}
	p.logger.Info("Finished scraping remote resource")
	return nil
}

func (p *Scraper) getPage() error {
	p.logger.Debug("Visiting page", "url", p.Url)
	httpResp, err := http.Get(p.Url)
	if err != nil {
		p.logger.Error("Scraping failed", "Error", err)
		return err
	}
	defer httpResp.Body.Close()
	if httpResp.StatusCode != http.StatusOK {
		p.logger.Error("Cannot visit "+p.Url, "Status code", httpResp.StatusCode, "Status", httpResp.Status)
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
	//TODO: clean up countries i.e. entries like (Portugal including islands)
	for list, s := range p.Selectors {
		var countries Country.Countries
		doc.Find(s).Each(func(i int, s *goquery.Selection) {
			countries = append(countries, Country.Country{
				Name: s.Text(),
			})
		})

		switch list {
		case "red":
			p.TravelLists.Red = countries
		case "amber":
			p.TravelLists.Amber = countries
		case "green":
			p.TravelLists.Green = countries
		}
	}
	return nil
}
