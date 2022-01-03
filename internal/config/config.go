package config

type (
	Config struct {
		Jwt      Jwt      `yaml:"jwt_generator"`
		Api      Api      `yaml:"api"`
		Database Database `yaml:"database"`
		Log      Log      `yaml:"log"`
	}

	Jwt struct {
		Secret string `yaml:"secret"`
	}

	Api struct {
		Port string `yaml:"port"`
	}

	Database struct {
		Host          string `yaml:"host"`
		Port          string `yaml:"port"`
		AuthDatabase  string `yaml:"auth_database"`
		AuthMechanism string `yaml:"auth_mechanism"`
		Username      string `yaml:"username"`
		Password      string `yaml:"password"`
	}

	Log struct {
		CmdLogPath  string `yaml:"cmd_log_path"`
		ApiLogPath  string `yaml:"api_log_path"`
		TestLogPath string `yaml:"test_log_path"`
	}
)
