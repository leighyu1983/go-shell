package controllers

import (
	"utils"
	"net/http"
	"entities"
	"github.com/gin-gonic/gin"
	"fmt"
)


// 创建单品信息
func updateHostName(c *gin.Context) {
	defer util.PanicHttpHandler(c)

	var hostname entities.HostnameEntity
	err := c.BindJSON(&hostname)
	if(err != nil) {
		panic(err)
	}

	
}