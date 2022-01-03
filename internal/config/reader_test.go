package config

import "testing"

func TestGetConfig(t *testing.T) {
	_, err := GetConfig("../../config.yaml")

	if err != nil {
		t.Errorf("Error loading enviroment variables: %s", err)
	}
}
