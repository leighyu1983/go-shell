package controllers

import (
	"fmt"
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"services"
)

// GET
func Terminal(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.Terminal(c.PostForm("ip"), c.PostForm("command"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": msg, "message": nil})
}