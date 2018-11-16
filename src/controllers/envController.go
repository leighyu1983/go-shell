package controllers

import (
	"fmt"
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"services"
)

// init
func Init(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	// target machine's ip
	msg := service.Init(c.PostForm("ip"))
	c.String(http.StatusOK, msg)
}

// GET
func Terminal(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.Terminal(c.PostForm("ip"), c.PostForm("command"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": msg, "message": nil})
}

// Post JSON
func SetHostName(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var hostname entities.HostnameEntity
	err := c.BindJSON(&hostname)
	if(err != nil) {
		panic(err)
	}

	fmt.Println("=========getting hostnmae params==========")
	msg, err1 := service.SetHostname(hostname.IP, hostname.Hostname)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": msg, "message": err1})
}

