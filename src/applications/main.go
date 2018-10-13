package main

import (
	"fmt"
	"utils"
	"daos"
 )

func main() {
	fmt.Println("started...");
	
	configObj, _ := util.GetConfig()
	fmt.Println(configObj.NetCard)
	daos.Test()
	
	fmt.Print("finished");
}

