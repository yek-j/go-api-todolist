package connector

import (
	"context"
	"encoding/json"
	"os"

	"to-do-list/ent"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Dbuser string `json:"dbuser"`
	DbServer string `json:"dbserver"`
	Dbname string `json:"dbname"`
	Dbpwd string `json:"dbpwd"`
}

func Connector() (*ent.Client, context.Context) {
	var config DBConfig 
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	
	file, err := os.Open(cwd + "\\db\\config.json")

	defer file.Close()
	if err != nil {
		panic(err)
	}

	configParser := json.NewDecoder(file)
	dcerr := configParser.Decode(&config)
	if dcerr != nil {
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