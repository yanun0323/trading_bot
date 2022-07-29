package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Init(cfgName string) {
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

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s ", err))
	}
}
