package main

import (
	"fmt"
	"utils"
	"daos"
 )

func main() {
	//test()

	
}

func test() {
	fmt.Println("test started...");
	
	configObj, _ := util.GetConfig()
	fmt.Println(configObj.NetCard)
	daos.Test()
	
	fmt.Print("test finished");
}
