package main

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Domain    string `yaml:"domain"`
	DebugMode bool   `yaml:"debugMode"`
}

var config Config

func surveyUser() {
	form := huh.NewForm(huh.NewGroup(
		huh.NewSelect[bool]().Title("Do you want Debug Mode Enabled (recommended)").Options(
			huh.NewOption("Yes", true),
			huh.NewOption("No", false),
		).Value(&config.DebugMode),
	),
		huh.NewGroup(
			huh.NewSelect[string]().Title("What domain are you in?").Options(
				huh.NewOption("us.nelnet.biz", ".us.nelnet.biz"),
				huh.NewOption("glhec.org", ".glhec.org"),
				huh.NewOption("nulsc.biz", ".nulsc.biz"),
			).Value(&config.Domain),
		),
	)
	form.Run()
}

func generateConfig() error {
	configPath := os.Getenv("XDG_CONFIG_HOME")

	if configPath == "" {
		homeDir, _ := os.UserHomeDir()
		configPath = homeDir + "/.config"
		err := os.MkdirAll(homeDir+".config", 0755)
		if err != nil {
			log.Error("Unable to create ~/.config")
		}
	}

	configPath += "/sheath-conf.yaml"

	if _, err := os.Stat(configPath); err != nil {
		err = writeConfig(configPath)
		return err
	}

	return nil
}

// TODO this should rewrite the config
func writeConfig(configFilePath string) error {
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Error("", err, "YAML Marshal Err")
	}

	err = os.WriteFile(configFilePath, conf, 0644)

	if err != nil {
		log.Fatal("", err, "Error Writing Config File to "+configFilePath)
	}

	return nil
}

func main() {
	log.SetReportTimestamp(false)
	log.SetReportCaller(false)

	surveyUser()
	e := generateConfig()

	if e != nil {
		log.Error("", e, "config file error")
	}

	log.Print(config)
	log.Print("Config Generated!")
}
