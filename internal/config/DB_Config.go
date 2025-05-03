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
	user, password, host, port, name string
}

func (dc *DatabaseConfiguration) LoadDBConfig() {
	err := godotenv.Load("../internal/config/env/.env")
	if err != nil {
		fmt.Println("Configuration (.env) file not found, switch to manual configuration ")
		dc.manualConfigurationSetting()
	} else {
		dc.user = os.Getenv("DB_USER")
		dc.password = os.Getenv("DB_PASSWORD")
		dc.host = os.Getenv("DB_HOST")
		dc.port = os.Getenv("DB_PORT")
		dc.name = os.Getenv("DB_NAME")
	}
}

func (dc *DatabaseConfiguration) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dc.user, dc.password, dc.host, dc.port, dc.name)
}

func (dc *DatabaseConfiguration) manualConfigurationSetting() {
	EnterText := map[*string]string{
		&dc.user: ui.ENTER_DB_USERNAME,
		&dc.host: ui.ENTER_DB_HOSTNAME,
		&dc.port: ui.ENTER_DB_PORT,
		&dc.name: ui.ENTER_DB_NAME,
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
		if dc.isPasswordValid(string(dbPassBytes)) {
			fmt.Println("Password cannot be empty")
			continue
		}
		dc.password = string(dbPassBytes)
		if dc.areRequiredFieldsFilled() {
			fmt.Println("\n", "Manual Configuration is successfully")
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
	return !(strings.Contains(dc.host, " ") || strings.Contains(dc.name, " "))
}
func (dc *DatabaseConfiguration) isPortNumberValid() bool {
	_, err := strconv.Atoi(dc.port)
	return err != nil
}
func (dc *DatabaseConfiguration) areRequiredFieldsFilled() bool {
	return dc.user != "" && dc.host != "" && dc.port != "" && dc.name != "" && dc.password != ""
}

func (dc *DatabaseConfiguration) isPasswordValid(pass string) bool {
	if strings.TrimSpace(pass) == "" {
		return true
	}
	return false
}
