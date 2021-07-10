package Settings

import (
	"github.com/hashicorp/go-hclog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configurations struct {
	Scraper struct {
		ServerName        string `yaml:"service-name"`
		Url               string `yaml:"url"`
		RedListSelector   string `yaml:"red-selector"`
		AmberListSelector string `yaml:"amber-selector"`
		GreenListSelector string `yaml:"green-selector"`
	} `yaml:"Scraper"`

	Global struct {
		LogLevel string       `yaml:"log-level"`
		logger   hclog.Logger `yaml:"-"`
	}`yaml:"Global"`
}

func Load(filename string, logger hclog.Logger) (*Configurations, error) {
	logger.Info("Loading config file", "config", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		logger.Error("Cannot load config file"+filename, "error", err)
		return nil,err
	}
	c := &Configurations{}
	yaml.Unmarshal(data, c)
	return c, nil;
}
