package main

type Conf struct {
	Redis RedisConf `yaml:"redis"`
	MySQL MySQLConf `yaml:"mysql"`
}

type RedisConf struct {
	Port int `yaml:"port"`
}

type MySQLConf struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}
