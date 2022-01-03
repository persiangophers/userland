package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

func GetConfig(path string) (*Config, error) {
	cfg := new(Config)
	err := readYAML(path, cfg)
	if err != nil {
		return &Config{}, err
	}

	return cfg, nil
}

func readEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

func readYAML(path string, cfg *Config) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if e := file.Close(); err == nil {
			err = e
		}
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return err
	}

	return nil
}
