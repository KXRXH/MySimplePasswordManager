package parser

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Theme  string `json:"theme"`
	DbPath string `json:"database_path"`
}

func ParseJson(path string) Config {
	content, err := ioutil.ReadFile(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		return Config{Theme: "StyleDefault", DbPath: "./pm"}
	}

	var cfg Config
	json.Unmarshal(content, &cfg)
	return cfg
}
