package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func (this *AppConfig) sampleConfigStr() string {
	return `# app config file (sample)
# you should follow the installation guide to configure the following fields
# (IMPORTANT) after your editing, the file name should be renamed to "config.yaml"

db_config:
  protocol: "tcp"            # Connection protocol, "tcp" is recommended and default
  host: "127.0.0.1"          # Database server host, e.g. "127.0.0.1" or "example.com"
  port: 3306                 # Database server port, default is 3306 for MySQL and MariaDB
  username: "root"           # Database username
  password: "change_me"      # Database password
  database: "app_db"         # Database name to connect
  table_prefix: "x_"         # Every data table prefix

redis_config:
  host: "127.0.0.1"          # Redis server host, e.g. "127.0.0.1" or "example.com"
  port: 6379                 # Redis server port, default is 6379
  username: ""               # Redis username (usually leave empty)
  password: ""               # Redis password (leave empty if not required)
  database: 0                # Redis database index (default is 0)

runtime:
  mode: "debug"                   # Mode of the Gin application, e.g. "debug" "test" "release"
  jwt_secret: "change_me"         # JWT secret key, must be changed to a random secure string
  access_token_expired_in: 3600   # Access token expiration time in seconds (e.g. 3600 = 1 hour)
  enable_register: true           # If set to "false", will not allow new user register
  min_password_len: 6             # Set the min length of user password allowance
  bcrypt_cost: 16                 # the bcrypt hash of the password at the given cost
  listening_port: "8080"          # Application listening port, e.g. "8080"`
}

var ExistingAppConfig AppConfig

func (this *AppConfig) ReadConfigFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("[ERROR] config file not found, creating new one: %v\n", err)
		this.CreateSampleConfigFile("./config/config_sample.yaml")
		os.Exit(1) // exit after creating sample file
	}
	if err := yaml.Unmarshal(data, this); err != nil {
		panic(fmt.Sprintf("[ERROR] failed to parse yaml: %v", err))
	}
}

func (this *AppConfig) CreateSampleConfigFile(path string) {
	// Write the predefined sample config string (with comments and default values)
	if err := os.WriteFile(path, []byte(this.sampleConfigStr()), 0644); err != nil {
		panic(fmt.Sprintf("[ERROR] failed to create sample config file: %v", err))
	}

	fmt.Printf("[INFO] sample config file created at %s, please edit it following the README.md\n", path)
}
