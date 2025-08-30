package config

type AppConfig struct {
	DbConfig    DbConfig    `yaml:"db_config" json:"db_config"`
	RedisConfig RedisConfig `yaml:"redis_config" json:"redis_config"`
	Runtime     Runtime     `yaml:"runtime"`
}

type DbConfig struct {
	Protocol    string `yaml:"protocol" json:"protocol"`
	Host        string `yaml:"host" json:"host"`
	Port        int    `yaml:"port" json:"port"`
	Username    string `yaml:"username" json:"username"`
	Password    string `yaml:"password" json:"password"`
	Database    string `yaml:"database" json:"database"`
	TablePrefix string `yaml:"table_prefix" json:"table_prefix"`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type Runtime struct {
	Mode                 string `yaml:"mode" yaml:"mode"`
	JwtSecret            string `yaml:"jwt_secret" json:"jwt_secret"`
	AccessTokenExpiredIn int    `yaml:"access_token_expired_in" json:"access_token_expired_in"`
	EnableRegister       bool   `yaml:"enable_register" json:"enable_register"`
	MinPasswordLen       int    `yaml:"min_password_len" json:"min_password_len"`
	BcryptCost           int    `yaml:"bcrypt_cost" json:"bcrypt_cost"`
	ListeningPort        string `yaml:"listening_port" json:"listening_port"`
}
