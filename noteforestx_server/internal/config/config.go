package config

type AppConfig struct {
	DbConfig     DbConfig     `yaml:"db_config" json:"db_config"`
	RedisConfig  RedisConfig  `yaml:"redis_config" json:"redis_config"`
	Runtime      Runtime      `yaml:"runtime"`
	Illustration Illustration `yaml:"illustration" json:"illustration"`
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
	Enabled  bool   `yaml:"enabled" json:"enabled"`
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
	ListeningAddr        string `yaml:"listening_addr" json:"listening_addr"`
}

type Illustration struct {
	SaveDir               string `yaml:"save_dir" json:"save_dir"`
	CompressedSmallPixel  int    `yaml:"compressed_small_pixel" json:"compressed_small_pixel"`
	CompressedMediumPixel int    `yaml:"compressed_medium_pixel" json:"compressed_medium_pixel"`
}
