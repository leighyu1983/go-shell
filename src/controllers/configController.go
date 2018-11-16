package controllers

import (
	// "fmt"
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"services"
)

/*****************************************
  Setup config for linux
 *****************************************/

// Post form
func SyncSshKey(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.SyncSshKey(c.PostForm("ip"), c.PostForm("password"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": msg})
}

// Post JSON
func SetHostName(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var hostname entities.HostnameEntity
	err := c.BindJSON(&hostname)
	if(err != nil) {
		panic(err)
	}

	msg := service.SetHostname(hostname.IP, hostname.Hostname)
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": msg})
}

// Post form
func DisbleFirewall(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.DisbleFirewall(c.PostForm("ip"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": msg})
}

// Post form
func DisbleSelinux(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.DisbleSelinux(c.PostForm("ip"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": msg})
}

// Post form
func SetTimezone(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	msg := service.SetTimezone(c.PostForm("ip"))
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": nil, "message": msg})
}