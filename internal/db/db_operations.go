package db

import (
	"LeetCode_Solutions/internal/structures"
	"context"
	"fmt"
)

func GetTasks(ctx context.Context) ([]structures.Task, error) {
	var tasks []structures.Task
	rows, err := Pool.Query(ctx, "SELECT task_id,task_number,task_name,task_description, level_id FROM tasks")
	if err != nil {
		fmt.Println("Database error")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if rows.Err() != nil {
			return tasks, rows.Err()
		}
		var taskItem structures.Task
		err = rows.Scan(&taskItem.TaskID, &taskItem.TaskNumber, &taskItem.TaskName, &taskItem.TaskDescription, &taskItem.LevelID)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, taskItem)
	}
	return tasks, err
}
