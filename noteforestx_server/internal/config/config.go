package config

type AppConfig struct {
	DbConfig    DbConfig    `yaml:"db_config" json:"db_config"`
	RedisConfig RedisConfig `yaml:"redis_config" json:"redis_config"`
	WxApp       WxApp       `yaml:"wx_app"`
	Runtime     Runtime     `yaml:"runtime"`
}

type DbConfig struct {
	Protocol string `yaml:"protocol" json:"protocol"`
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database string `yaml:"database" json:"database"`
}

type RedisConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     int    `yaml:"port" json:"port"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	Database int    `yaml:"database" json:"database"`
}

type WxApp struct {
	AppId     string `yaml:"app_id"`
	AppSecret string `yaml:"app_secret"`
}

type Runtime struct {
	JwtSecret            string `yaml:"jwt_secret"`
	AccessTokenExpiredIn int    `yaml:"access_token_expired_in"`
	ListeningPort        string `yaml:"listening_port"`
}
