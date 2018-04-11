package utils

import (
	"fmt"
	"os/exec"
)

func CreateObject(path string) {
	//CmdCreate := exec.Command("kubectl", "create", "-f", path + "/" + name)
	CmdCreate := exec.Command("kubectl", "create", "-f", path)
	fmt.Println("kubectl " + "create " + "-f " + path)
	//CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin")
	CmdOut, err := CmdCreate.Output()
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl create")
	//fmt.Println(string(CmdOut))

}