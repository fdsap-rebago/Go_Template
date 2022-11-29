// Package connect provides ...
package middleware

import (
	"Template/middleware/go-utils/database"
	"Template/middleware/go-utils/encryptDecrypt"
	httpUtils "Template/middleware/go-utils/http"
	"fmt"

	"log"
	"net/http"
)

func CreateConnection() {
	username, dbConfigErr := encryptDecrypt.Decrypt(GetEnv("POSTGRES_USERNAME"), GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	password, dbConfigErr := encryptDecrypt.Decrypt(GetEnv("POSTGRES_PASSWORD"), GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	host, dbConfigErr := encryptDecrypt.Decrypt(GetEnv("POSTGRES_HOST"), GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	dbName, dbConfigErr := encryptDecrypt.Decrypt(GetEnv("DATABASE_NAME"), GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	fmt.Println("username: ", username)
	fmt.Println("password: ", password)
	fmt.Println("host: ", host)
	fmt.Println("dbName: ", dbName)
	httpUtils.Client.New(&http.Client{})
	database.PostgreSQLConnect(
		username,
		password,
		host,
		dbName,
		GetEnv("POSTGRES_PORT"),
		GetEnv("POSTGRES_SSL_MODE"),
		GetEnv("POSTGRES_TIMEZONE"),
	)
	err := database.DBConn.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}

}
