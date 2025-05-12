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

func GetTaskByNumber(ctx context.Context, num int) (structures.Task, error) {
	var task structures.Task
	row := Pool.QueryRow(ctx,
		"SELECT task_id,task_number,task_name,task_description, level_id FROM tasks WHERE task_number=$1", num)
	err := row.Scan(&task.TaskID, &task.TaskNumber, &task.TaskName, &task.TaskDescription, &task.LevelID)
	return task, err
}

func CreateTask(ctx context.Context, number, levelID int, name, description string) error {
	query := "INSERT INTO tasks (task_number, task_name, task_description, level_id) VALUES ($1, $2,$3,$4)"
	_, err := Pool.Exec(ctx, query, number, name, description, levelID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTaskByNumber(ctx context.Context, task structures.Task) error {
	query := `UPDATE tasks SET task_number = $1, task_name = $2, task_description = $3, 
                 level_id = $4, WHERE task_id = $5;`
	_, err := Pool.Exec(ctx, query,
		task.TaskNumber,
		task.TaskName,
		task.TaskDescription,
		task.LevelID,
		task.TaskID,
	)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}
	return nil
}
