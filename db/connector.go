package connector

import (
	"context"
	"log"

	"to-do-list/ent"

	_ "github.com/go-sql-driver/mysql"
)

func Connector() (*ent.Client, context.Context) {
	db, err := ent.Open("mysql", "root:MYSQLroot1234@tcp(localhost:3306)/todolist?parseTime=True")
	if err != nil {
		log.Fatalf("DB 접속 실패 mysql 버전: %v", err)
	} 
	
	ctx := context.Background()

	
	

	return db, ctx;
}