package configuration

import (
	"flag"
)

var (
	Cfg *config

	configPath string
)

// Initialize initializes package.
func Initialize() {
	//nolint:exhaustruct
	Cfg = &config{}

	flag.StringVar(&configPath, "config", "", "Path to configuration file")
}
