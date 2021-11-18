package main

import (
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/viper"
	"github.com/thanhpp/aws-cognito-example/infra/cognitolib"
	"github.com/thanhpp/aws-cognito-example/infra/httpserver"
)

var (
	name = "config.yml"
)

func main() {
	cfg := readConfig(name)
	if cfg == nil {
		panic("nil config")
	}

	cognitoApp, err := cognitolib.SetupCognito(cfg.Region, cfg.PoolID, cfg.AppID)
	if err != nil {
		panic(err)
	}

	server := httpserver.NewServer(cfg.Host, cfg.Port, cognitoApp)
	go func() {
		if err := server.Start(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	<-sigChan

	if err := server.Stop(); err != nil {
		panic(err)
	}
}

func readConfig(name string) *Config {
	var cfg = new(Config)

	file := filepath.Base(name)

	// 0: filename, 1: file ext
	fileSl := strings.Split(file, ".")
	if len(fileSl) != 2 {
		return nil
	}

	v := viper.New()
	v.SetConfigName(fileSl[0])
	v.SetConfigType(fileSl[1])
	v.AddConfigPath(".")     // current folder
	v.AddConfigPath("..")    // depth 1 - test file
	v.AddConfigPath("../..") // depth 2 - test file

	if err := v.ReadInConfig(); err != nil {
		return nil
	}

	if err := v.Unmarshal(cfg); err != nil {
		return nil
	}

	return cfg
}
