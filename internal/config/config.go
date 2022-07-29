package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init(cfgName string) error {
	envCfgName := os.Getenv("CONFIG_NAME")
	if len(envCfgName) != 0 {
		cfgName = envCfgName
	}
	if len(cfgName) == 0 {
		cfgName = "config"
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName(cfgName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./internal/config")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error config file: %w ", err)
	}
	return nil
}
