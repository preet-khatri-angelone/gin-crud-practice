package controllers

import (
	"CRUD-GIN/model"
	"CRUD-GIN/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error" : "invalid input",
		})
		ctx.Abort()
	}

	if ok := utils.CreateUser(user); !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "could not create user",
		})
		ctx.Abort()
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : err,
		})
		ctx.Abort()
	}

	ctx.JSON(http.StatusOK, gin.H{"token: ": token})
}

func Login(ctx *gin.Context) {
	user := &model.User{}
	if err := ctx.ShouldBindJSON(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		ctx.Abort()
	}

	if _, ok := utils.FetchUser(user); !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "could not fetch user",
		})
		ctx.Abort()
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		fmt.Println("error: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error" : "error in generating token",
		})
		ctx.Abort()
	}

	ctx.JSON(200, gin.H{
		"token: " : token,
	})
}

