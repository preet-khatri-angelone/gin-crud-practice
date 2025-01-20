package controllers

import (
	"CRUD-GIN/model"
	"CRUD-GIN/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create(ctx *gin.Context)  {
	task := &model.Task{}
	if err := ctx.ShouldBindJSON(task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error: " : "invalid input",
		})
		return
	}

	userID := ctx.Param("userid")

	userid, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("could not convert userID to int")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in creating task",
		})
		return
	}

	if ok := utils.InsertTask(task, userid); !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in creating task",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message" : "task successfully created",
	})
}

func Task(ctx *gin.Context) {
	taskID := ctx.Param("taskid")
	taskid, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("could not convert taskID to int")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in fetching task",
		})
		return
	}

	task, ok := utils.FetchTask(taskid)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in fetching task",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Task: " : task,
	})
}

func Tasks(ctx *gin.Context) {
	userID := ctx.Param("userid")
	userid, err := strconv.Atoi(userID)
	if err != nil {
		fmt.Println("could not convert userID to int")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in fetching tasks",
		})
		return
	}

	tasks, ok := utils.FetchTasks(userid)
	if !ok {
		fmt.Println("error in fetching tasks")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in fetching tasks",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Tasks: " : tasks,
	})
}

func Delete(ctx *gin.Context) {
	taskID := ctx.Param("taskid")
	taskid, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("could not convert taskID to int")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in deleting task",
		})
		return
	}
	if ok := utils.DeleteTask(taskid); !ok {
		fmt.Println("could not delete")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in deleting task",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "task deleted successfully",
	})
}

func Update(ctx *gin.Context)  {
	newTask := model.Task{}
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error: " : "invalid input",
		})
		return
	}

	taskID := ctx.Param("taskid")
	taskid, err := strconv.Atoi(taskID)
	if err != nil {
		fmt.Println("could not convert taskID to int")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in updating task",
		})
		return
	}

	if ok:= utils.UpdateTask(taskid, newTask); !ok {
		fmt.Println("error in updating the task")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in updating task",
		})
		return
	}

	updatedTask, ok := utils.FetchTask(taskid)
	if !ok {
		fmt.Println("error in fetching task id")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message" : "error in updating task",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"update Task: " : updatedTask,
	})
}
