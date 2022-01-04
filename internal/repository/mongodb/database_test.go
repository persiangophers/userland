package mongodb

import "testing"

func TestDropDatabase(t *testing.T) {

	client, err := InitClient(cfg, logger)

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	_, err = client.NewDatabase("users")

	if err != nil {
		t.Errorf("Error initializing mongodb client: %s", err)
	}

	err = client.DisconnectDatabase()

	if err != nil {
		t.Errorf("Error diconnecting mongodb client: %s", err)
	}
}
