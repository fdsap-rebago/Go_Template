// Package connect provides ...
package middleware

import (
	"Template/middleware/go-utils/database"
	httpUtils "Template/middleware/go-utils/http"

	"log"
	"net/http"
)

func CreateConnection() {

	httpUtils.Client.New(&http.Client{})
	database.PostgreSQLConnect(
		GetEnv("POSTGRES_USERNAME"),
		GetEnv("POSTGRES_PASSWORD"),
		GetEnv("POSTGRES_HOST"),
		GetEnv("DATABASE_NAME"),
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_SSL_MODE"),
		GetEnv("POSTGRES_TIMEZONE"),
	)
	err := database.DBConn.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}

}
