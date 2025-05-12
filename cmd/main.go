package main

import (
	"LeetCode_Solutions/internal/config"
	"LeetCode_Solutions/internal/db"
	"context"
	"fmt"
)

func main() {
	dbc := config.DatabaseConfiguration{}
	dbc.LoadDBConfig()
	err := db.Connect(dbc)
	if err != nil {
		return
	}
	ctx := context.Background()
	task, err := db.GetTaskByNumber(ctx, 9)
	fmt.Println(task)
}
