package configa

import (
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"gopkg.in/yaml.v3"
)

// # for something to be configurable it must have gen and read METHODS

// i have no idea how set options as yaml or how to Unmarshal into a struct here
type Options struct {
	//yaml tags somehow
	Opts map[string]any
}

type Config struct {
	BaseName string
	Path     string
	// embed a struct for options
	Options Options
}

type Configurable interface {
	GenerateConfig() (Config, error)
	ReadConfig() (Config, error)
	CreateConfigPath()
}

// This should set options but idk how to handle the options struct
func (c *Config) SurveyUser(form huh.Form) {
	form.Run()
}

func (c *Config) SetOptions(options Options) {
	c.Options = options
}

// Default implementation places a config file directly in xdg home
func (c *Config) SetConfigPath() error {

	configDirPath := os.Getenv("XDG_CONFIG_HOME")

	if configDirPath == "" {
		homeDir, _ := os.UserHomeDir()
		configDirPath = homeDir + "/.config"
		err := os.MkdirAll(homeDir+".config", 0755)
		if err != nil {
			log.Error("Unable to create ~/.config")
			return err
		}
	}

	// Set ConfigPath
	c.Path = configDirPath + c.BaseName
	return nil

}

// This function should always generate and overwrite a config.
func (c *Config) GenerateConfig() error {
	conf, err := yaml.Marshal(&config)

	if err != nil {
		log.Error("", err, "YAML Marshal Err")
	}

	err = os.WriteFile(c.Path, conf, 0644)

	if err != nil {
		log.Fatal("", err, "Error Writing Config File to "+c.Path)
		return err
	}

	return nil
}

func (c *Config) ReadConfig() {
	f, err := os.ReadFile(c.Path)

	if err != nil {
		log.Fatal(err)
	}

	// Check the pointer here
	// if err := yaml.Unmarshal(f, &c); err != nil {
	if err := yaml.Unmarshal(f, c); err != nil {
		log.Fatal(err)
	}

	//TODO: Set c.Options here
}

func NewConfig(configBaseName string) Config {
	c := Config{BaseName: configBaseName}
	err := c.SetConfigPath()
	if err != nil {
		log.Error(err)
	}

	return c
}
