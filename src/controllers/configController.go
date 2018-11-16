package controllers

import (
	"fmt"
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"services"
)

/*****************************************
  Setup config for linux
 *****************************************/

func SyncSSHKeyFile(c *gin.Context) {
	defer util.PanicHttpHandler(c)
	
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

