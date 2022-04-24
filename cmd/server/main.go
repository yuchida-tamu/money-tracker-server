package main

import (
	"fmt"

	"github.com/yuchida-tamu/money-tracker-server/internal/db"
	"github.com/yuchida-tamu/money-tracker-server/internal/record"
	transportHttp "github.com/yuchida-tamu/money-tracker-server/internal/transport/http"
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

	httpHandler := transportHttp.NewHandler(recordService)
	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main(){
	fmt.Println("Go REST")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}