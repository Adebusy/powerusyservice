package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func (database *Database) LoadDVInstance() (db *sql.DB, err error) {
	godotenv.Load()
	configReader, readErr := godotenv.Read()
	if readErr != nil {
		fmt.Println("Unable to read .env file Error occurred " + readErr.Error())
		return
	}
	var fff = configReader["Pasaword"]

	fmt.Println(fff)

	serverConn := fmt.Sprintf("server=%s; database=%s;integrated security=true; User ID=%s; Password=%s", configReader["DatabaseServer"], configReader["Database"], configReader["UserID"], configReader["Password"])
	db, err = sql.Open("sqlserver", serverConn)

	if err != nil {
		//database.Logger.Error("Error connecting to DB", err.Error())
		fmt.Println(err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		//database.Logger.Error("Error connecting to DB", err.Error())
		fmt.Println(err.Error())
	}
	database.DB = db
	//database.Logger.Info("Database connection was successful")
	return db, err
}

type Database struct {
	DB *sql.DB
}
