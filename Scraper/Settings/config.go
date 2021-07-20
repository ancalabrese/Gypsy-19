package Settings

import (
	"io/ioutil"

	"github.com/ancalabrese/Gypsy-19/Scraper/Data"
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v2"
)

type Configurations struct {
	Scraper struct {
		ServerName        string            `yaml:"service-name"`
		Url               string            `yaml:"url"`
		HtmlListSelectors map[string]string `yaml:"html-selectors"`
	} `yaml:"Scraper"`

	Global struct {
		LogLevel string       `yaml:"log-level"`
		logger   hclog.Logger `yaml:"-"`
	} `yaml:"Global"`

	Connector struct {
		Repo          Data.Repo       `yaml:"repo"`
		CommitterInfo Data.CommitInfo `yaml:"committer-info"`
	} `yaml:"github"`
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
