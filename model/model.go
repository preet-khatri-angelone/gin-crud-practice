package model

type User struct {
	ID       int    `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	TaskID   int    `json:"taskID"`
	UserID   int    `json:"userID"`
	TaskName string `json:"taskName"`
	Status   bool   `json:"status"`
}
