package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var (
	App      = new(app)
	Database = new(database)
	Redis    = new(redis)
	Log      = new(logger)
	COS      = new(cos)
)

type app struct {
	Name                   string
	Port                   string
	Env                    string
	TokenSecret            string
	TokenEffectiveDuration uint16
}

type database struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
	Charset  string
}

type redis struct {
	Host     string
	Port     uint16
	Secret   string
	Database int8
}

type logger struct {
	Filename string
	Level    string
}

type cos struct {
	Bucket    string
	Region    string
	SecretID  string
	SecretKey string
}

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("加载配置文件失败：%v", err)
	}

	if err := cfg.Section("app").MapTo(App); err != nil {
		log.Fatalf("app配置映射失败：%v", err)
	}

	if err := cfg.Section("database").MapTo(Database); err != nil {
		log.Fatalf("database配置映射失败：%v", err)
	}

	if err := cfg.Section("redis").MapTo(Redis); err != nil {
		log.Fatalf("redis配置映射失败：%v", err)
	}

	if err := cfg.Section("log").MapTo(Log); err != nil {
		log.Fatalf("log配置映射失败：%v", err)
	}

	if err := cfg.Section("cos").MapTo(COS); err != nil {
		log.Fatalf("cos配置映射失败：%v", err)
	}
}
