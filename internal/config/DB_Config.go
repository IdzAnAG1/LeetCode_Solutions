package config

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var (
	errInput   error
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
)

func LoadDBConfig() {
	err := godotenv.Load("../internal/config/env/.env")
	if err != nil {
		fmt.Println("Configuration (.env) file not found, switch to manual configuration ")
	} else {
		dbUser = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
		dbName = os.Getenv("DB_NAME")
	}
}

func GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbUser, dbPassword,
		dbHost, dbPort, dbName)
}

func manualConfigurationSetting() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Database Username : ")
	dbUser, errInput = reader.ReadString('\n')
	if errInput != nil {

	}
	fmt.Print("Enter Database Password : ")
	dbPassword, errInput = reader.ReadString('\n')
	if errInput != nil {

	}
}
