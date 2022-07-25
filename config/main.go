package config

type dev struct {
	Host string
	Port int16
}

type mysql struct {
	User     string
	Password string
	Address  string
	Schema   string
	Port     int16
}

type redis struct {
}

type Config struct {
	Dev   *dev
	Mysql *mysql
	Redis *redis
}

func InitializeConfig() *Config {
	config := &Config{
		Dev: &dev{
			Host: "127.0.0.1",
			Port: 7001,
		},
		Mysql: &mysql{
			User:     "root",
			Password: "",
			Address:  "127.0.0.1:3306",
			Schema:   "go_tiny_code",
		},
		Redis: &redis{},
	}

	return config
}
