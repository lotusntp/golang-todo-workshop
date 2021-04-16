package main

import (
	"test/go/api/todo"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()


	r.GET("/",todo.GetAll)


	r.Run()

}