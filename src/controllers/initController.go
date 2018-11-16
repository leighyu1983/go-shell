package controllers

import (
	"utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"services"
)

/*****************************************
 Setup linux environment at the beginning 
 *****************************************/


// init master machine
func Init(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	// target machine's ip
	msg := service.PrepareEnv()
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": msg, "message": nil})
}

