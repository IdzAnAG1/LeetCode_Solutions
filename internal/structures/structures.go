package structures

type LeetcodeTask struct {
	TaskID          int    `json:"task_id"`
	TaskNumber      int    `json:"task_number"`
	TaskName        string `json:"task_name"`
	TaskDescription string `json:"task_description"`
	Level           int    `json:"level"`
}
type Level struct {
	LevelID         int `json:"level_id"`
	DifficultyLevel int `json:"difficulty_level"`
}
type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type Solution struct {
	SolutionID   int    `json:"solution_id"`
	TaskID       int    `json:"task_id"`
	LanguageID   int    `json:"language_id"`
	SolutionText string `json:"solution_text"`
	Complexity   string `json:"complexity"`
	Updated      string `json:"updated_at"`
	Created      string `json:"created_at"`
}
