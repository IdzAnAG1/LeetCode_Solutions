package main

import (
	"LeetCode_Solutions/internal/config"
	"fmt"
)

func main() {
	db := config.DatabaseConfiguration{}
	db.LoadDBConfig()

	fmt.Println(db.URL())
}
