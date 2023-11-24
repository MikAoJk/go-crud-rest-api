package router

import (
	"github.com/MikAoJk/go-crud-rest-api/src/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/users", postUsers)
	r.GET("/users/:id", getUser)
	r.GET("/users", getUsers)
	r.PUT("/users/", updateUser)
	r.DELETE("/users/:id", deleteUser)
	return r
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

func updateUser(ctx *gin.Context) {
	var updatedUser db.User
	err := ctx.Bind(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbUser, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbUser.Name = updatedUser.Name
	dbUser.Email = updatedUser.Email

	res, err := db.UpdateUser(dbUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func getUsers(ctx *gin.Context) {
	res, err := db.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": res,
	})
}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func postUsers(ctx *gin.Context) {
	var user db.User
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": res,
	})
}
