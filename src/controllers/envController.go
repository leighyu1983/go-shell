package controllers

import (
	"fmt"
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"services"
)


func SetHostName(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var hostname entities.HostnameEntity
	err := c.BindJSON(&hostname)
	if(err != nil) {
		panic(err)
	}

	fmt.Println("=========getting hostnmae params==========")
	fmt.Println(hostname);
	msg, err1 := service.SetHostname(hostname.IP, hostname.Hostname)
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": msg, "message": err1})
}

