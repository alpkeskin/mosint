/*
Copyright © 2023 github.com/alpkeskin
*/
package config

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Services struct {
		BreachDirectoryApiKey string `yaml:"breach_directory_api_key"`
		EmailRepApiKey        string `yaml:"emailrep_api_key"`
		HunterApiKey          string `yaml:"hunter_api_key"`
		IntelXApiKey          string `yaml:"intelx_api_key"`
		HaveIBeenPwnedApiKey  string `yaml:"haveibeenpwned_api_key"`
	}
	Settings struct {
		IntelXMaxResults int `yaml:"intelx_max_results"`
	}
}

var (
	Cfg *Config
)

func New() *Config {
	return &Config{}
}

func (c *Config) Parse(cfgFile string) error {
	path := getConfigPath(cfgFile)
	data, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	var cFile Config
	err = yaml.Unmarshal(data, &cFile)

	if err != nil {
		return err
	}

	Cfg = &cFile

	return nil
}

func (c *Config) Exists(cfgFile string) bool {
	path := getConfigPath(cfgFile)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func getConfigPath(cfgFile string) string {

	if !strings.EqualFold(cfgFile, "") {
		return cfgFile
	}

	homeDir, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/.mosint.yaml", homeDir)
}
