package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetAll(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"messgae":"ok"})
}