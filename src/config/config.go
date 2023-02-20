package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
)

func Load() {
	fmt.Println("Environment variables load started.")

	error := godotenv.Load()
	if error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("SERVER_PORT"))
	if error != nil {
		log.Fatal(error)
	}

	ConnectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_NAME"),
	)

	fmt.Println("Environment variables load completed.")
}
