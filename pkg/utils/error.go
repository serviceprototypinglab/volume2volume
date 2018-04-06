package utils

import "fmt"

func checkErrorMessage(err error, message string){
	if err != nil {

		fmt.Println(message)
		//panic(err)
		fmt.Println(err)
		fmt.Println("Error managed by checkErrorMessage")
	}
}
