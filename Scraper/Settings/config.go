package Settings

import (
	"io/ioutil"
	"time"

	"github.com/ancalabrese/Gypsy-19/Scraper/DB/GitHubClient"
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v2"
)

type Configurations struct {
	Scraper struct {
		ServerName       string `yaml:"service-name"`
		Url              string `yaml:"url"`
		HtmlListSelector string `yaml:"html-selector"`
	} `yaml:"scraper"`

	Global struct {
		LogLevel string       `yaml:"log-level"`
		logger   hclog.Logger `yaml:"-"`
	} `yaml:"global"`

	Connector struct {
		Repo          GitHubClient.Repo       `yaml:"repo"`
		CommitterInfo GitHubClient.CommitInfo `yaml:"committer-info"`
	} `yaml:"github"`

	Service struct {
		Name               string        `yaml:"name"`
		Url                string        `yaml:"url"`
		Port               string        `yaml:"port"`
		ApiBasePath        string        `yaml:"api-base-path"`
		CorsAllowedOrigins string        `yaml:"cors-allowed-origins"`
		ApiRoutes          []string      `yaml:"api-routes"`
		IdleTimeout        time.Duration `yaml:"idle_timeout"`
		ReadTimeout        time.Duration `yaml:"read-timeout"`
		WriteTimeout       time.Duration `yaml:"write-timeout"`
	} `yaml:"server"`
}

func Load(filename string, logger hclog.Logger) (*Configurations, error) {
	logger.Info("Loading config file", "config", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("Cannot load config file"+filename, "error", err)
		return nil, err
	}
	c := &Configurations{}
	yaml.Unmarshal(data, c)
	return c, nil
}
