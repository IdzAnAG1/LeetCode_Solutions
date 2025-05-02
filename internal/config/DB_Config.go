package config

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
)

type DatabaseConfiguration struct {
	User, Password, Host, Port, Name string
}

var (
	errInput error // Убрать глобальную переменную !!!
)

func (dc *DatabaseConfiguration) LoadDBConfig() {
	err := godotenv.Load("../internal/config/env/.env")
	if err != nil {
		fmt.Println("Configuration (.env) file not found, switch to manual configuration ")
		dc.manualConfigurationSetting()
	} else {
		dc.User = os.Getenv("DB_USER")
		dc.Password = os.Getenv("DB_PASSWORD")
		dc.Host = os.Getenv("DB_HOST")
		dc.Port = os.Getenv("DB_PORT")
		dc.Name = os.Getenv("DB_NAME")
	}
}

func (dc *DatabaseConfiguration) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dc.User, dc.Password,
		dc.Host, dc.Port, dc.Name)
}

func (dc *DatabaseConfiguration) manualConfigurationSetting() {
	var dbPassBytes []byte
	fmt.Print("Enter Database Username : ")
	readFromCLI(&dc.User)
	fmt.Print("Enter Database Password : ")
	dbPassBytes, errInput = term.ReadPassword(int(syscall.Stdin))
	if errInput != nil {
		fmt.Printf("Error at enter password: %v\n", errInput)
	}
	dc.Password = string(dbPassBytes)
	fmt.Print("\nEnter Database Host name : ")
	readFromCLI(&dc.Host)
	fmt.Print("Enter Database Port Number : ")
	readFromCLI(&dc.Port)
	fmt.Print("Enter Database Name : ")
	readFromCLI(&dc.Name)
}

func readFromCLI(field *string) {
	reader := bufio.NewReader(os.Stdin)
	*field, errInput = reader.ReadString('\n')
	if errInput != nil {
		fmt.Printf("An error occurred when entering data :%v \n", errInput)
	}
	*field = strings.TrimSpace(*field)
}
