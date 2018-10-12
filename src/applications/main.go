package main

import (
	"fmt"
	"utils"
 )

func main() {
	fmt.Print("started...");
	
	configObj, _ := util.GetConfig()
	fmt.Println(configObj.NetCard)

	fmt.Print("finished");
}

