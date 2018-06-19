package utils

import (
	"fmt"
	"os/exec"
)

// login in the openshift cluster
func LoginCluster(cluster, username, password string) {
	username = "--username=" + username
	password = "--password=" + password
	CmdLogin := exec.Command("oc", "login", cluster, username, password, "--insecure-skip-tls-verify=true")
	//CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin")
	CmdOut, err := CmdLogin.Output()
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running login")
	fmt.Println(string(CmdOut))
}

// Login to the user admin (minishift)
// TODO chante it for production clusters
func LoginAdmin(cluster string) {

	CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin", "--insecure-skip-tls-verify=true")
	//CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin")
	CmdOut, err := CmdLogin.Output()
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running login")
}

// Get the objects of the typeObjects in the cluster and the project that you login beforehand.
func GetObjects(typeObject string) string {
	CmdGetDeployments := exec.Command("oc", "get", typeObject, "-o", "json")
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("getObjects error in type " + typeObject)
		return ""
	}
	//checkErrorMessage(err, "Error running get " + typeObject)
	return string(CmdOut)
}

// Change to the project: projectName.
func ChangeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	CheckErrorMessage(err, "Error running: change project")
	fmt.Println(string(CmdProjectOut))
}