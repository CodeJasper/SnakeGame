package config

import (
	"context"
	"log"
	"os"

	// "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var Connect *pgx.Conn

var dbUser string
var dbPass string
var dbHost string
var dbName string

func ConfigDB() {
	getEnvironmentVariables()
	config, err := pgx.ParseConfig("postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":26257/" + dbName + "?sslmode=require")
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}
	// Connect to the database.
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	Connect = conn

}

func getEnvironmentVariables() {
	err := godotenv.Load("dbSettings.env")

	if err != nil {
		log.Fatalln("Ha ocurrido un error al buscar el archivo de las variables de entorno ", err)
	}

	dbUserEnv, definedDbUser := os.LookupEnv("DB_USER")
	dbHostEnv, definedDbHost := os.LookupEnv("DB_HOST")
	dbPassEnv, definedDbPass := os.LookupEnv("DB_PASS")
	dbNameEnv, definedDbName := os.LookupEnv("DB_NAME")

	if !definedDbUser || !definedDbHost || !definedDbPass || !definedDbName {
		log.Fatalln("Alguna de las variables de entorno requerida no esta definida")
	}

	dbUser = dbUserEnv
	dbPass = dbPassEnv
	dbHost = dbHostEnv
	dbName = dbNameEnv

}
