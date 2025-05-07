package structures

import "time"

type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}
type Language struct {
	LanguageID   int    `json:"language_id"`
	LanguageName string `json:"language_name"`
}
type Level struct {
	LevelID         int `json:"level_id"`
	DifficultyLevel int `json:"difficulty_level"`
}
type Solution struct {
	SolutionID   int       `json:"solution_id"`
	TaskID       int       `json:"task_id"`
	LanguageID   int       `json:"language_id"`
	SolutionText string    `json:"solution_text"`
	Complexity   string    `json:"complexity"`
	Updated      time.Time `json:"update_at"`
	Created      time.Time `json:"created_at"`
}
type TaskCategory struct {
	TaskID     int `json:"task_id"`
	CategoryID int `json:"category_id"`
}
type Task struct {
	TaskID          int    `json:"task_id"`
	TaskNumber      int    `json:"task_number"`
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	LevelID         int    `json:"level_id"`
}
