package Twitter

import (
	"context"
	"fmt"
	"strings"

	"github.com/ancalabrese/Gypsy-19/Scraper/Data/Country"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/hashicorp/go-hclog"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/array"
	"golang.org/x/oauth2"
	"google.golang.org/api/transport/http"
)

type Twitter struct {
	Enabled        bool   `yaml:"enabled"`
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
	AccessKey      string `yaml:"access_key"`
	AccessSecret   string `yaml:"access_secret"`
	logger         hclog.Logger
	client         *twitter.Client
}

func (t *Twitter) Init(l hclog.Logger) {
	t.logger = l
	config := oauth1.NewConfig(t.ConsumerKey, t.ConsumerSecret)
	token := oauth1.NewToken(t.AccessKey, t.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	t.client = twitter.NewClient(httpClient)

}

func (t *Twitter) OnUpdate(newList Country.Lists, oldList Country.Lists) {
	var msg strings.Builder
	var diff int
	tags := []string{"#Covid19", "#TravelList", "#Update", "#TravelRestrictions"}

	if diff = len(newList.Red) - len(oldList.Red); diff > 0 {
		msg.WriteString(fmt.Sprintf("+%d countries moved to Red list", diff))
	}
	if diff = len(newList.Amber) - len(oldList.Amber); diff > 0 {
		msg.WriteString(fmt.Sprintf("\n+%d countries moved to Amber list", diff))
	}
	if diff = len(newList.Green) - len(oldList.Green); diff > 0 {
		msg.WriteString(fmt.Sprintf("\n+%d countries moved to Green list", diff))
	}

	msg.WriteString(fmt.Sprintf("Check out the Covid Travel Restriction Tracker -> \"https://github.com/ancalabrese/COVID-Travel-Restriction-Tracker\""))
	msg.WriteString(strings.Join(tags, " "))
	if err := t.postUdpate(msg.String()); err != nil {
		t.logger.Error("Couldn't post tweet", "error", err)
	}
}

func (t *Twitter) postUdpate(msg string) error {
	tweet, _, err := t.client.Statuses.Update(msg, nil)
	if err == nil {
		t.logger.Debug("Tweeted", "tweet", tweet)
	}
	return err
}
