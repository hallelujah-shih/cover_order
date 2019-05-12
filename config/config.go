package config

import (
	"github.com/alecthomas/gometalinter/_linters/src/gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type DBConfig struct {
	Type string `json:"type" yaml:"type"`
	Url  string `json:"url" yaml:"url"`
}

type ConverFile struct {
	FileList []string `json:"file_list" yaml:"file_list"`
}

type Config struct {
	DB    *DBConfig   `json:"db" yaml:"db"`
	Files *ConverFile `json:"files" yaml:"files"`
}

func New(fpath string) (*Config, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	datas, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(datas, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
