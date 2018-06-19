package utils

import "fmt"


// functions that show the error in the console with some message
func CheckErrorMessage(err error, message string){
	if err != nil {

		fmt.Println(message)
		//panic(err)
		fmt.Println(err)
		fmt.Println("Error managed by checkErrorMessage")
	}
}
