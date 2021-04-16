package main

import (
	"net/http"
	"test/go/api/todo"
	"test/go/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	database.ConnectDatabase()


	r.GET("/",todo.GetAll)
	r.GET("/user",todo.GetTodoList)
	r.POST("/",todo.CreateTotoLists)
	r.DELETE("/delete/:id",todo.DeleteList)
	r.POST("/upload",todo.Upload)
	r.StaticFS("/file",http.Dir("public"))
	

	r.Run()

}