/*
 * @Author: fengzhilaoling fengzhilaoling@gmail.com
 * @Date: 2025-11-29 09:40:26
 * @LastEditors: fengzhilaoling
 * @LastEditTime: 2025-11-29 13:53:00
 * @FilePath: \ginManager\config\config.go
 * @Description: 配置文件读取
 * Copyright (c) 2025 by fengzhilaoling@gmail.com, All Rights Reserved.
 */
package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
	DB     DB     `mapstructure:"database"`
	Log    Log    `mapstructure:"log"`
	JWT    JWT    `mapstructure:"jwt"`
}

type Server struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DB struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	Path       string `mapstructure:"path"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type JWT struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

var C = new(Config)

// Init 读取并解析
func Init(cfgPath string) {
	if cfgPath == "" {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
	} else {
		viper.SetConfigFile(cfgPath)
	}

	if err := viper.ReadInConfig(); err != nil {
		panic("cannot read config: " + err.Error())
	}
	if err := viper.Unmarshal(C); err != nil {
		panic("cannot unmarshal config: " + err.Error())
	}

	// 确保日志目录存在
	_ = os.MkdirAll(C.Log.Path, 0755)
}
