package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	Reels          map[string][][]string     `json:"reels,omitempty"`           // key is a reel name, value is an array of reels, reel is an array of symbols
	Weights        map[string]map[string]int `json:"weights,omitempty"`         // key is a weight name, value is a map where key is a variant and value is its weight
	Symbols        map[string]string         `json:"symbols,omitempty"`         // key is a symbol name, value is a symbol code (as on reels)
	LinePaytable   map[string][]int          `json:"line_paytable,omitempty"`   // key is a symbol code, value is an array where index is a count of symbols and value is a payout
	SymbolPaytable map[string]map[string]int `json:"count_paytable,omitempty"`  // key is a symbol code, value is a map where key is a count and value is a bet multiplier
	FreespinConfig FreespinConfig            `json:"freespins"`                 // complex structure that contains attributes related to Freespin
	PurchasePrices map[string]int            `json:"purchase_prices,omitempty"` // key is a purchase name, value is multiplier that applies to bet

	// Set of storages to keep any value during rounds and all game iterations
	StorageInt    map[string]int
	StorageBool   map[string]bool
	StorageString map[string]string
}

type FreespinConfig struct {
	TriggerSymbol  string         `json:"trigger_symbol,omitempty"`
	TriggerCount   map[string]int `json:"trigger_count,omitempty"`
	RetriggerCount map[string]int `json:"retrigger_count,omitempty"` // may be omitted if game doesn't have retrigger or uses its own logic for retriggers
}

func NewConfig() *Config {
	c := Config{
		StorageInt:    map[string]int{},
		StorageBool:   map[string]bool{},
		StorageString: map[string]string{},
	}

	return &c
}

func ReadConfig(path string, config *Config) {
	data, err := ioutil.ReadFile(fmt.Sprintf("pkg/%s", path))
	if err != nil {
		log.Fatal("unable to read json config file: ", err)
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal("unable to unmarshall json config: ", err)
	}
}
