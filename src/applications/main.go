package main

import (
	"fmt"
	"utils"
 )

func main() {
	fmt.Println("started...");
	
	configObj, _ := util.GetConfig()
	fmt.Println(configObj.NetCard)

	fmt.Print("finished");
}

