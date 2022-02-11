package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件错误, err", err)
		return
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("Debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func LoadData(file *ini.File) {
	DBHost = file.Section("database").Key("DBHost").MustString("127.0.0.1")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassword = file.Section("database").Key("DBPassword").MustString("123456")
	DBName = file.Section("database").Key("DBName").MustString("ginblog")
}
