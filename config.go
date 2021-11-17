package main

type Config struct {
	Host      string `mapstructure:"Host"`
	Port      string `mapstructure:"Port"`
	Region    string `mapstructure:"Region"`
	PoolID    string `mapstructure:"PoolID"`
	AppID     string `mapstructure:"AppID"`
	AppSecret string `mapstructure:"AppSecret"`
}
