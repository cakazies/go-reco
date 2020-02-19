package configs

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"github.com/spf13/viper"
)

var (
	// GetDB variable for connection DB
	GetDB *gorm.DB // connection about DB
)

func Connect() {
	host := viper.GetString("configDB.host")
	port := viper.GetString("configDB.port")
	user := viper.GetString("configDB.user")
	password := viper.GetString("configDB.password")
	dbname := viper.GetString("configDB.dbname")

	conn := fmt.Sprintf("postgresql://%s%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	revel.AppLog.Debugf(conn)
	db, err := gorm.Open("postgres", conn)

	if err != nil {
		log.Println(err)
	}
	// defer db.Close()
	GetDB = db
}

func StartViper() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		revel.AppLog.Debugf("error viper", err)
	}
}
