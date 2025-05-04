package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Laps        int    `json:"laps"`
	LapLen      int    `json:"lapLen"`
	PenaltyLen  int    `json:"penaltyLen"`
	FiringLines int    `json:"firingLines"`
	Start       string `json:"start"`
	StartDelta  string `json:"startDelta"`
}

func DownloadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	config := Config{}

	if err = json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("decode file: %w", err)
	}

	return &config, nil
}
