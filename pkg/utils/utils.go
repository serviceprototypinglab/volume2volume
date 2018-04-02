package utils

import (
	"io/ioutil"
	"fmt"
	"encoding/json"
)


//Json functions
func ReadJson(path, name string) map[string]interface{} {
	fmt.Println(path)
	plan, _ := ioutil.ReadFile(path + "/" + name + ".json")
	//fmt.Println(plan)
	//var data []interface{}
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error reading json")
		//fmt.Println(data)
		fmt.Println(err)
	}
	return data
}


func ReadJsonData(path string) []map[string]interface{} {
	fmt.Println(path)
	plan, _ := ioutil.ReadFile(path + "/data.json")
	//fmt.Println(plan)
	//var data []interface{}
	var data []map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error reading json")
		//fmt.Println(data)
		fmt.Println(err)
	}
	return data
}