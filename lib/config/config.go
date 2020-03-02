package config

import (
	"gopkg.in/ini.v1"
)

var Config ConfigList

type ConfigList struct {
	CategoryName string
	Port         int
	DbName       string
	SQLDriver    string
	Length       int
	NumberLength int
	SymbolLength int
}

func init() {
	cfg, _ := ini.Load("config/application.ini")
	Config = ConfigList{
		Length:       cfg.Section("app").Key("default_passwd_length").MustInt(16),
		NumberLength: cfg.Section("app").Key("default_number_length").MustInt(2),
		SymbolLength: cfg.Section("app").Key("default_symbol_length").MustInt(2),
		CategoryName: cfg.Section("app").Key("category_name").String(),
		Port:         cfg.Section("web").Key("port").MustInt(8888),
		DbName:       cfg.Section("db").Key("name").String(),
		SQLDriver:    cfg.Section("db").Key("driver").String(),
	}
}
