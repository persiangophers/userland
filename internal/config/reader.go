package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// ReadEnv reads some configs from environment variables
func ReadEnv(cfg *Config) error {
	return envconfig.Process("", cfg)
}

func ReadYAML(path string, cfg *Config) (err error) {
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
