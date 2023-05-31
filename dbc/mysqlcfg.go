package dbc

import (
	"encoding/json"
	"fmt"
	"os"
)

type psqlCfg struct {
	User         string `json:"user"`
	Pass         string `json:"pass"`
	Host         string `json:"host"`
	Port         string `json:"port"`
	DatabaseName string `json:"db_name"`
}

func configure(path string) (*psqlCfg, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can't open config file: %s", err)
	}
	defer f.Close()
	cfg := &psqlCfg{}
	jsonParser := json.NewDecoder(f)
	err = jsonParser.Decode(cfg)
	if err != nil {
		return nil, fmt.Errorf("parse config failed: %s", err)
	}
	return cfg, nil
}
