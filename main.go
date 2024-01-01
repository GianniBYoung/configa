package main

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	domain    string `yaml:"domain"`
	debugMode bool   `yaml:"debugMode"`
}

var config Config

func surveyUser() {
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[bool]().Title("Do you want Debug Mode Enabled (recommended)").Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).Value(&config.debugMode),
	),
		huh.NewGroup(
			huh.NewSelect[string]().Title("What domain are you in?").Options(
				huh.NewOption("test.com", "test.com"),
				huh.NewOption("test2.com", "test2.com"),
			).Value(&config.domain),
		),
	)
	form.Run()
}

func main() {

	log := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
	})

	surveyUser()

	log.Print("", config, "Configuration Struct")
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal("", err, "error")
	}

	log.Print(string(conf))

	// output:
	// {test.com true}="Configuration Struct"
	// {}
}
