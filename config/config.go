package config

import (
	"fmt"
	"github.com/spf13/viper"
	"payment/manager"
	"payment/util"
)

type ApiConfig struct {
	Url          string
	AppName      string
	SignatureKey string
}

type Manager struct {
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

type DbConfig struct {
	Host     string
	User     string
	Port     string
	Password string
	Name     string
}

type Config struct {
	Manager
	DbConfig
	ApiConfig
	LogLevel string
}

func (c Config) Configuration(path, fileName string) Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	c.DbConfig = DbConfig{
		Host:     v.GetString("db.db_host"),
		Port:     v.GetString("db.db_port"),
		User:     v.GetString("db.db_user"),
		Password: v.GetString("db.db_password"),
		Name:     v.GetString("db.db_name"),
	}

	c.ApiConfig = ApiConfig{
		Url:          v.GetString("api.url"),
		AppName:      v.GetString("api.app_name"),
		SignatureKey: v.GetString("api.signature_key"),
	}

	c.LogLevel = v.GetString("api.log_level")

	return c
}

func NewConfig(path, name string) Config {
	cfg := Config{}
	cfg = cfg.Configuration(path, name)
	util.NewLog(cfg.LogLevel)

	dataSourceName := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)
	cfg.InfraManager = manager.NewInfraManager(dataSourceName)
	cfg.RepoManager = manager.NewRepoManager(cfg.InfraManager)
	cfg.UseCaseManager = manager.NewUseCaseManager(cfg.RepoManager)

	return cfg
}
