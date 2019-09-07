package configuration

import (
	// stdlib
	"flag"
)

var (
	Cfg *config

	configPath string
)

// Initialize initializes package.
func Initialize() {
	Cfg = &config{}

	flag.StringVar(&configPath, "config", "", "Path to configuration file")
}
