package config

var Conf Config

type Config struct {

	MysqlConfig Mysql
}


func NewConfig(){

	var config Config
	config.MysqlConfig.InitMysqlConfig()

	Conf = config
}
