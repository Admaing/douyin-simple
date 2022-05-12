package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//将ini文件配置参数初始化

var (
	// Debug
	AppMode  string
	HttpPort string
	JwtKey   string

	//	Db
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("ini文件读取失败", err)
	}
	LoadServer(file)
	LoadDatabase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}
func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("root")
	DbUser = file.Section("database").Key("DbUser").MustString("12345678990")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

}
