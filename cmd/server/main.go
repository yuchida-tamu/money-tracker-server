package main

import (
	"context"
	"fmt"

	"github.com/yuchida-tamu/money-tracker-server/internal/db"
	"github.com/yuchida-tamu/money-tracker-server/internal/record"
)

// Run - is going to be responsible for
// the instantiation and startup of out go application
func Run() error {
	fmt.Println("starting up our application")

	// try to create a new database instance
	db, err := db.NewDatabase()
	if err != nil{
		fmt.Println("failed to connect to the database")
		return err
	}
	
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	recordService := record.NewService(db)

	recordService.PostRecord(
		context.Background(),
		record.Record{
			ID: "6ba00572-c32a-11ec-9d64-0242ac120002",
			DATE_CREATED: "2022-4-26",
			AMOUNT: 1200,
			CATEGORY: "food",
			RECORD_DESCRIPTION: "ramen restaurant",
			EXPENSE_TYPE: "expense",
		},
	)

	fmt.Println(recordService.GetRecord(
		context.Background(),
		"6ba00572-c32a-11ec-9d64-0242ac120002",
	))

	return nil
}

func main(){
	fmt.Println("Go REST")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}