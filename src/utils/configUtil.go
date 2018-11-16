package util

import(
	"entities"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

func GetConfig()(entities.ConfigEntity, error) {
	data, err := ioutil.ReadFile("./config.yaml")
	fmt.Println("......configuration data:\n" + string(data))
	
	config := entities.ConfigEntity{}
    yaml.Unmarshal(data, &config)
	return config, err
}
