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

func surveyUser() Config {

	var (
		domain    string
		debugMode bool
	)

	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[bool]().Title("Do you want Debug Mode Enabled (recommended)").Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).Value(&debugMode),
	),
		huh.NewGroup(
			huh.NewSelect[string]().Title("What domain are you in?").Options(
				huh.NewOption("us.nelnet.biz", ".us.nelnet.biz"),
				huh.NewOption("us.glhec.org", ".us.glhec.org"),
				huh.NewOption("nulsc.biz", ".nulsc.biz"),
			).Value(&domain),
		),
	)
	form.Run()

	return Config{
		domain:    domain,
		debugMode: debugMode,
	}

}

func main() {

	log := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    false,
		ReportTimestamp: false,
	})

	config := surveyUser()

	log.Print("", config, "Configuration")
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Fatal("", err, "error")
	}

	log.Print(string(conf))
	f, err := os.Create("config.yaml")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(conf))
	if err != nil {
		panic(err)
	}
}
