package connector

import (
	"context"
	"to-do-list/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Dbuser string `json:"dbuser"`
	DbServer string `json:"dbserver"`
	Dbname string `json:"dbname"`
	Dbpwd string `json:"dbpwd"`
}

func Connector() (*ent.Client, context.Context) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
	
	var config DBConfig
	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	mysqlConfig := config.Dbuser+":"+config.Dbpwd+"@tcp("+config.DbServer+")/"+config.Dbname+"?parseTime=True"

	db, err := ent.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	} 
	
	ctx := context.Background()

	return db, ctx;
}