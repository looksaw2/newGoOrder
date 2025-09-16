package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)



func NewViperConfig() error {
	dir ,err := os.Getwd()
	if err != nil {
		return err
	}
	configPath := filepath.Join(dir,"../common/config")
	//设置读取global.yaml文件
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	viper.EnvKeyReplacer(strings.NewReplacer("_","-"))
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}