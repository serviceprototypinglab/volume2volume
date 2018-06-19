package utils

import (
	"fmt"
	"os/exec"
)

// Deploy the object in path
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

// Create a secret. All the info of the object should be in templates/secrets
// secretName supported are s3 or minio
func CreateSecret(secretName string) {
	auxSecretName := secretName
	if secretName == "s3" {
		auxSecretName = "AWS"
	}
	if secretName == "minio" {
		auxSecretName = "MINIO"
	}
	// TODO change the path to one from configuration
	path := "./templates/secrets/"
	CmdCreate := exec.Command("kubectl", "create", "secret", "generic", secretName + "-secret",
		"--from-file=" + path + "RESTIC_PASSWORD",
		"--from-file=" + path + auxSecretName + "_ACCESS_KEY_ID",
		"--from-file=" + path + auxSecretName + "_SECRET_ACCESS_KEY")
	//fmt.Println("kubectl " + "create " + "-f " + path)
	CmdOut, err := CmdCreate.Output()
	fmt.Println("OUT")
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl create generic secret " + path)
}


// Get size of the volume (used for stats)
func GetSizeVolume(podName, containerName, pathData string) string {
	CmdCreate := exec.Command("kubectl", "exec", podName, "--", "du", "-sh", pathData)
	CmdOut, err := CmdCreate.Output()
	fmt.Println(string(CmdOut))
	CheckErrorMessage(err, "Error running kubectl exec")
	return string(CmdOut)
	//kubectl exec arkismongopersistentd0-3083001275-lwn8w -- du -sh ./data/db
}