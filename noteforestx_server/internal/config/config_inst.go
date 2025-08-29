package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func sampleConfigStr() string {
	return `# app config file
# you should follow the installation guide to configure the following fields

db_config:
  protocol: "tcp"            # Connection protocol, "tcp" is recommended
  host: "127.0.0.1"          # Database server host
  port: 3306                 # Database server port, default is 3306 for MySQL
  username: "root"           # Database username
  password: "change_me"      # Database password
  database: "app_db"         # Database name to connect

redis_config:
  host: "127.0.0.1"          # Redis server host
  port: 6379                 # Redis server port, default is 6379
  username: ""               # Redis username (usually leave empty)
  password: ""               # Redis password (leave empty if not required)
  database: 0                # Redis database index (default is 0)

runtime:
  jwt_secret: "change_me"    # JWT secret key, must be changed to a random secure string
  access_token_expired_in: 3600   # Access token expiration time in seconds (e.g. 3600 = 1 hour)
  listening_port: "8080"     # Application listening port, e.g. "8080"`
}

var ExistingAppConfig AppConfig

func (this *AppConfig) ReadConfigFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[ERROR] config file not found, creating new one: %v\n", err)
		this.CreateSampleConfigFile(path)
		os.Exit(1) // exit after creating sample file
	}
	if err := yaml.Unmarshal(data, this); err != nil {
		panic(fmt.Sprintf("[ERROR] failed to parse yaml: %v", err))
	}
}

func (this *AppConfig) CreateSampleConfigFile(path string) {
	// Write the predefined sample config string (with comments and default values)
	if err := os.WriteFile(path, []byte(sampleConfigStr()), 0644); err != nil {
		panic(fmt.Sprintf("[ERROR] failed to create sample config file: %v", err))
	}

	fmt.Printf("[INFO] sample config file created at %s, please edit it following the README.md\n", path)
}
