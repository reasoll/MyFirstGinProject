package config

import _ "gopkg.in/yaml.v2"

type Mysql struct {
	Addr		 string	`yaml:"Addr"`
	Port		 string	`yaml:"Port"`
	UserNmae	 string	`yaml:"UserName"`
	Password	 string	`yaml:"Password"`
	DataBaseName string	`yaml:"DataBaseName"`
}

func (mysql *Mysql)DefaultMySqlConfig() {
	mysql.Addr = "172.16.1.66"
	mysql.Port = "3306"
	mysql.DataBaseName = "outsystems"
	mysql.UserNmae = "root"
	mysql.Password = "root"
}

func (mysql *Mysql) InitMysqlConfig() {
	mysql.DefaultMySqlConfig();
}
