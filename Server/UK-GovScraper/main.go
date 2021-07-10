package main

import (
	// "fmt"
	"github.com/hashicorp/go-hclog"
	settings "./Settings"
)

func main() {

	l := hclog.New(&hclog.LoggerOptions{
		Level: hclog.Debug,
	})

	//Load service config
	conf :=  settings.Create().Load("config.yml")
}
