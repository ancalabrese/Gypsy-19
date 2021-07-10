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
	} `yaml:"scraper"`

	Global struct {
		LogLevel string       `yaml:"log-level"`
		logger   hclog.Logger `yaml:"-"`
	}
}

func Create() *Configurations {
	return &Configurations{};
}

func (c *Configurations) Load(filename string) error {
	c.Global.logger.Info("Loading Config", "config", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		c.Global.logger.Error("Cannot load config file"+filename, "error", err)
		return err
	}
	yaml.Unmarshal(data, c)
	return nil
}
