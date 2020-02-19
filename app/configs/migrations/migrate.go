package main

import (
	"go-reco/app/configs"
	"go-reco/app/models"
	"log"

	"github.com/revel/revel"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./go-reco/conf/")
	err := viper.ReadInConfig()
	if err != nil {
		revel.AppLog.Debugf("error viper", err)
	}
}

// running go run go-reco/app/configs/migrations/migrate.go
func main() {
	log.Println("loading ...")
	configs.Connect()
	configs.GetDB.AutoMigrate(&models.Student{})
	log.Println("Migraion success ...")
	defer configs.GetDB.Close()
}
