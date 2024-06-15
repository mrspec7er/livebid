package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/mrspec7er/livebid/server/internal/database"
	"github.com/mrspec7er/livebid/server/internal/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	DBConn := database.StartConnection()

	config := &server.Config{
		DB: DBConn,
	}

	dbConn, err := DBConn.DB.DB()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
	defer dbConn.Close()

	server := server.NewInstance(config)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
