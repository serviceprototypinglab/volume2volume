package utils

import (
	"fmt"
	"io/ioutil"
	"os"
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

func WriteJson(path, name string, data map[string]interface{}) error {
	//write json in path restic
	f, err3 := os.Create(path +"/"+ name +".json")
	if err3 != nil {
		fmt.Println("Error creating data.json")
		fmt.Println(err3)
		return err3
	} else {
		objectOs, err2 := json.Marshal(data)
		if err2 != nil {
			fmt.Println("Error creating the json object")

			fmt.Println(err2)
			return err2
		} else {
			f.WriteString(string(objectOs))
			f.Sync()
			fmt.Println("Created  data.json in" + path )
		}
	}
	return nil
}

