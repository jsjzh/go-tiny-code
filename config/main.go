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

type Config struct {
	Dev   *dev
	Mysql *mysql
}

func InitializeConfig() *Config {
	config := &Config{
		Dev: &dev{
			Host: "127.0.0.1",
			Port: 7001,
		},
		Mysql: &mysql{
			User:     "souche_rw",
			Password: "Ehz74BrKfOIwA5Zg",
			Address:  "test.database3702.scsite.net:3702",
			Schema:   "wireless_devops_proto",
		},
	}

	return config
}
