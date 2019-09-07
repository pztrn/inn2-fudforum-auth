package configuration

import (
	// stdlib
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	// other
	"gopkg.in/yaml.v2"
)

// This structure represents application's configuration.
type config struct {
	Debug    bool `yaml:"debug"`
	Database struct {
		DSN        string `yaml:"dsn"`
		Parameters string `yaml:"parameters"`
		Prefix     string `yaml:"prefix"`
	} `yaml:"database"`
	Groups struct {
		Default string `yaml:"default"`
		Groups  []struct {
			Group string   `yaml:"group"`
			Users []string `yaml:"users"`
		}
	} `yaml:"groups"`
}

// Checks neccessary parameters for filling.
func (c *config) checkParameters() {
	if c.Database.DSN == "" {
		log.Fatalln("database/dsn parameter isn't filled, don't know to which database I should connect!")
	}
}

// Initialize initializes configuration structure by reading configuration
// file and populate structure with it's data.
func (c *config) Initialize() {
	// If -config wasn't provided or empty - do nothing.
	if configPath == "" {
		log.Fatalln("-config parameter is empty, don't know where configuration is!")
	}

	// Normalize configuration path.
	if strings.HasPrefix(configPath, "~") {
		userDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln("Failed to get user's home directory for user! Error was: " + err.Error())
		}

		configPath = strings.Replace(configPath, "~", userDir, 1)
	}

	absolutePath, err := filepath.Abs(configPath)
	if err != nil {
		log.Fatalln("Failed to get configuration file's absolute path! Error was: " + err.Error())
	}

	// Read and parse.
	dataAsBytes, err1 := ioutil.ReadFile(absolutePath)
	if err1 != nil {
		log.Fatalln("Failed to read configuration file! Error was: " + err1.Error())
	}

	err2 := yaml.Unmarshal(dataAsBytes, c)
	if err2 != nil {
		log.Fatalln("Failed to parse configuration file! Error was: " + err2.Error())
	}

	if c.Debug {
		log.Printf("Configuration parsed: %+v\n", c)
	}

	c.checkParameters()
}
