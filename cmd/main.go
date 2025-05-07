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
	ctx := context.Context(context.Background())
	tasks, err := db.GetTasks(ctx)
	if err != nil {
		fmt.Printf("err : %v", err)
	}
	for _, val := range tasks {
		res := fmt.Sprintf("\tTask_ID : %d,\n\tTask_Number : %d,\n\tTask_Name : %s,\n\tTask_Description : %s,\n\tLevel_ID : %d",
			val.TaskID, val.TaskNumber, val.TaskName, val.TaskDescription, val.LevelID)
		fmt.Println(res)
	}
}
