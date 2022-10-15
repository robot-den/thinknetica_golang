package model

import (
	"log"
	"strings"
)

// availableModes is a set of available modes
var availableModes = map[string]bool{
	"base":                true,
	"fs-buy":              true,
	"bonus-buy":           true,
	"fs-chance":           true,
	"bonus-chance":        true,
	"fs-and-bonus-change": true,
}

// ParseVariant splits variant string into mode and math version and sets them to game config
func ParseVariant(variant string, config *Config) {
	flags := strings.Split(variant, ":")
	if len(flags) != 2 || flags[1] == "" {
		log.Fatal("variant params is invalid, use --help to read about it")
	}

	config.StorageString["math_version"] = flags[1]
	mode := flags[0]

	if !availableModes[mode] {
		log.Fatal("invalid mode for game")
	}

	config.StorageBool[mode] = true
}
