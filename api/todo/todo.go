package todo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"test/go/database"

	"github.com/gin-gonic/gin"
)

type CreateTodo struct {
	Username string `json:"username" binding:"required"`
	Title string `json:"title" binding:"required"`
	Messgae string `json:"message" binding:"required"`
}

func GetAll(c *gin.Context){
	var todoList []database.Todo

	database.DB.Find(&todoList)
	c.JSON(http.StatusOK , gin.H{"Result": todoList})
}

func CreateTotoLists(c *gin.Context)  {
	var input CreateTodo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error})
	}

	todoList := database.Todo{Username: input.Username,Title: input.Title,Messgae:input.Messgae}

	database.DB.Create(&todoList)

	c.JSON(http.StatusOK,gin.H{"Data:": todoList})

	 
}

func GetTodoList(c *gin.Context)  {
	var todoList []database.Todo

	if err := database.DB.Where("username=?",c.Query("username")).Find(&todoList).Error;
	err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"List":"Record not found!"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"List":todoList})
}


func DeleteList(c *gin.Context)  {
	var todoList database.Todo
	
	if err := database.DB.Where("id = ?",c.Param("id")).First(&todoList).Error;

	err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error:":"Record not found!"})
	}

	database.DB.Delete(&todoList)

	c.JSON(http.StatusOK,gin.H{"data": true})
}


func Upload(c *gin.Context){
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprint("file err : %s",err.Error()))
		return
	}
	filename := header.Filename
	out,err:= os.Create("public/"+filename)
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
	_,err = io.Copy(out,file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/file/" + filename
	c.JSON(http.StatusOK,gin.H{"filepath:":filepath})
}