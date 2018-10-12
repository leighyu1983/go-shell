package main

import (
	"fmt"
	"utils"
 )

func main() {
	fmt.Print("started...");
	
	configObj, _ := util.GetConfig()
	fmt.Println(configObj.ImageUrlDish)

	fmt.Print("finished");
}

