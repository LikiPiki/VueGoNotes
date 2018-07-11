package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

var (
	Cnf Config
)

type Config struct {
	Port string `json:"PORT"`
  DB_PATH string `json:"DB"`
  DB string `json:"DB_Path"`
}

func ReadConfig() error {
  var err error
  var path string

  path, err = filepath.Abs("./config/config.json")

  var raw []byte
  raw, err = ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	json.Unmarshal(raw, &Cnf)
	return nil
}
