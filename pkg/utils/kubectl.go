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
	fmt.Println("OUT")
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl create -f " + path)
	//fmt.Println(string(CmdOut))

}

func CreateSecret(secret string) {
	path := "templates/secrets/"
	//CmdCreate := exec.Command("kubectl", "create", "-f", path + "/" + name)
	CmdCreate := exec.Command("kubectl", "create", "-f", path)
	fmt.Println("kubectl " + "create " + "-f " + path)
	//CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin")
	CmdOut, err := CmdCreate.Output()
	fmt.Println("OUT")
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl create -f " + path)
	//fmt.Println(string(CmdOut))

}



func GetSizeVolume(podName, containerName, pathData string) string {
	CmdCreate := exec.Command("kubectl", "exec", podName, "--", "du", "-sh", pathData)
	CmdOut, err := CmdCreate.Output()
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl exec")
	return string(CmdOut)
	//kubectl exec arkismongopersistentd0-3083001275-lwn8w -- du -sh ./data/db
}