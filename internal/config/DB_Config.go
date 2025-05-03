package config

import (
	"LeetCode_Solutions/internal/config/ui"
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/term"
	"os"
	"strconv"
	"strings"
	"syscall"
)

type DatabaseConfiguration struct {
	User, Password, Host, Port, Name string
}

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
	EnterText := map[*string]string{
		&dc.User: ui.ENTER_DB_USERNAME,
		&dc.Host: ui.ENTER_DB_HOSTNAME,
		&dc.Port: ui.ENTER_DB_PORT,
		&dc.Name: ui.ENTER_DB_NAME,
	}
	for {
		for field, prompt := range EnterText {
			dc.readFromCLI(field, prompt)
		}
		if !dc.isHostAndNameValid() {
			fmt.Println("Database Hostname or Database Name cannot contains \" \" in his title")
			continue
		}
		if dc.isPortNumberValid() {
			fmt.Println("The port number cannot be anything other than a number")
			continue
		}

		fmt.Print(ui.ENTER_DB_PASSWORD)
		dbPassBytes, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Printf(ui.ERROR_ENTERING_PASSWORD, err)
		}
		if dc.isPasswordValid(dbPassBytes) {
			fmt.Println("Password cannot be empty")
			continue
		}
		dc.Password = string(dbPassBytes)

		if dc.areRequiredFieldsFilled() {
			fmt.Println("Manual Configuration is successfully")
			break
		}
		fmt.Println(" Fields Database configuration cannot be empty\n Try again . . .\n ")
	}
}

func (dc *DatabaseConfiguration) readFromCLI(field *string, text string) {
	reader := bufio.NewReader(os.Stdin)
	var err error
	for {
		fmt.Print(text)
		*field, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf(ui.ERROR_ENTERING_DATA, err)
			continue
		}
		if *field != "" && strings.TrimSpace(*field) != "" {
			*field = strings.TrimSpace(*field)
			break
		}
		fmt.Println("This field cannot be empty")
	}
}

func (dc *DatabaseConfiguration) isHostAndNameValid() bool {
	return !(strings.Contains(dc.Host, " ") || strings.Contains(dc.Name, " "))
}
func (dc *DatabaseConfiguration) isPortNumberValid() bool {
	_, err := strconv.Atoi(dc.Port)
	return err != nil
}
func (dc *DatabaseConfiguration) areRequiredFieldsFilled() bool {
	return dc.User != "" && dc.Host != "" && dc.Port != "" && dc.Name != ""
}

func (dc *DatabaseConfiguration) isPasswordValid(pass []byte) bool {
	if strings.TrimSpace(string(pass)) == "" {
		return false
	}
	return true
}
