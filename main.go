package main

import (
	"io"
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
				huh.NewOption("us.nelnet.biz", ".us.nelnet.biz"),
				huh.NewOption("glhec.org", ".glhec.org"),
				huh.NewOption("nulsc.biz", ".nulsc.biz"),
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

	log.Print("", config, "Configuration")
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal("", err, "error")
	}

	log.Print(string(conf))
	f, err := os.Create("config.yaml")

	if err != nil {
		log.Fatal("", err, "Fatal Error")
	}
	defer f.Close()

	_, err = io.WriteString(f, string(conf))
	if err != nil {
		log.Fatal("", err, "Fatal Error")
	}
}
