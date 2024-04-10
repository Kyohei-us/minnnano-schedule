package infra

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLConnector struct {
	Conn *gorm.DB
}

func NewPostgreSQLConnector() *PostgreSQLConnector {
	dsn, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		fmt.Println("ERROR: DB Connection String not found")
		panic("ERROR: DB Connection String not found")
	}

	db_connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		fmt.Println("DB connected!")
	} else {
		fmt.Println("ERROR: DB Connection failed")
		panic(err)
	}

	return &PostgreSQLConnector{
		Conn: db_connection,
	}
}
