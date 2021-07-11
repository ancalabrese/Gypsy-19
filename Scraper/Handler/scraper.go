package Handler

import (
	"errors"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/ancalabrese/Gypsy-19/Scraper/Settings"
	"github.com/hashicorp/go-hclog"
)

type Country struct {
	Name string
	Note string
}

type Scraper struct {
	url         string
	selectors   map[string]string
	travelLists map[string][]Country
	logger      hclog.Logger
}

func CreateScraper(config *Settings.Configurations, l hclog.Logger) *Scraper {
	s := &Scraper{
		url:       config.Scraper.Url,
		selectors: config.Scraper.HtmlListSelectors,
		travelLists: make(map[string][]Country),
		logger:    l,
	}
	s.logger.Debug("Created new scraper", "url", s.url)
	return s
}

func (p *Scraper) RetrieveLists() {
	p.logger.Info("Starting scraping remote resource")
	p.getPage()
	p.logger.Info("Finished scraping remote resource")
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
	p.logger.Debug("Retrieved countries", "Lists", p.travelLists)
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
	for list, s := range p.selectors {
		doc.Find(s).Each(func(i int, s *goquery.Selection) {
			country := Country{
				Name: s.Text(),
			}
			p.travelLists[list] = append(p.travelLists[list], country)
		})
	}
	return nil
}
