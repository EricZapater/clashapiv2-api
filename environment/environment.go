package environment

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Site is the data structure that represents an Site.
type Environment struct {
	Bearer     string
	DbHost     string
	DbUser     string
	DbPass     string
	DbPort     int
	DbName     string
	Clan string
}

// LoadEnvironment creates a new Environment.
func LoadEnvironment() Environment {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	godotenv.Load(".env")

	// getting env variables SITE_TITLE and DB_HOST
	bearer := os.Getenv("BEARER")
	dbHost := os.Getenv("DBHOST")
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbPort, err := strconv.Atoi(os.Getenv("DBPORT"))
	CheckConversionError("DBPORT", err)
	dbName := os.Getenv("DBNAME")
	clan := os.Getenv("CLAN")

	return Environment{
		Bearer:     bearer,
		DbHost:     dbHost,
		DbUser:     dbUser,
		DbPass:     dbPass,
		DbPort:     dbPort,
		DbName:     dbName,
		Clan: clan,
	}
}

func CheckConversionError(varName string, err error) {
	if err != nil {
		log.Fatalf("%s environment variable not found", varName)
	}
}
