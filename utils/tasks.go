package utils

import (
	"CRUD-GIN/db"
	"CRUD-GIN/model"
	"fmt"
)

func InsertTask(task *model.Task, userID int) bool {
	_, err := db.DB.Exec(`INSERT INTO tasks(userid, taskname, status) VALUES ($1, $2, $3)`, userID, task.TaskName, task.Status)
	if err != nil {
		fmt.Println("error in creating task: ", err)
		return false
	}
	return true
}

func FetchTask(taskID int) (*model.Task, bool) {
	task := &model.Task{}
	row := db.DB.QueryRow(`SELECT * FROM tasks WHERE taskid = $1`, taskID)
	if err := row.Scan(&task.UserID, &task.TaskID, &task.TaskName, &task.Status); err != nil {
		fmt.Println("error in fetching task: ", err)
		return nil, false
	}
	return task, true
}

func FetchTasks(userID int) ([]model.Task, bool) {
	rows, err := db.DB.Query(`SELECT * FROM tasks WHERE userid = $1`, userID)
	if err != nil {
		fmt.Println("error in executing query: ", err)
		return nil, false
	}

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.TaskID, &task.UserID, &task.TaskName, &task.Status); err != nil {
			fmt.Println("error in reading the rows", err)
			return nil, false
		}
		tasks = append(tasks, task)
	}

	return tasks, true
}

func DeleteTask(taskID int) bool {
	if _, err := db.DB.Exec(`DELETE FROM tasks WHERE taskid = $1`, taskID); err != nil {
		fmt.Println("error in deleting task: ", err)
		return false
	}
	return true
}

func UpdateTask(taskID int, updatedTask model.Task) bool {
	res, err := db.DB.Exec(`UPDATE tasks SET taskname = $1, status = $2 WHERE taskid = $3`, updatedTask.TaskName, updatedTask.Status, taskID)
	if err != nil {
		fmt.Println("error in updating the task: ", err)
		return false
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("error in affecting rows: ", err)
		return false
	}

	if rowsAffected == 0 {
		fmt.Println("no rows affected ", err)
		return false
	}
	return true
}
