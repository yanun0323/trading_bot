package main

import (
	"fmt"
	"main/internal/app"
	"main/internal/config"

	"github.com/spf13/viper"
)

func main() {
	if err := config.Init(""); err != nil {
		panic(err)
	}

	if err := app.NewApp(viper.GetString("setting.symbol")); err != nil {
		fmt.Println(err)
	}
}
