package app

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
	"github.com/mitchellh/go-homedir"
	"volume2volume/pkg/utils"
)

func Example() {
	fmt.Println("example")
}

func PairsVolumesByName(PathData, PathTemplate, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo,
	UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) ([]map[string]interface{},
	[]map[string]interface{})  {
	//Read  Volumes/ClusterFrom/data.json
	var from [] map[string]interface{}
	var to [] map[string]interface{}
	GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
	clusterFromVolumes := utils.ReadJsonData(PathData + "/ClusterFrom")
	clusterToVolumes := utils.ReadJsonData(PathData + "/ClusterTo")
	fmt.Println("read it")
	for _,v := range clusterFromVolumes {
		for _,k := range clusterToVolumes {
			if v["deploymentName"] == k["deploymentName"] {
				if v["volumeName"] == k["volumeName"] {
					fmt.Println(v["volumeName"])
					from = append(from, v)
					to = append(to, k)
				}
			}
		}
	}
	return from, to
}

//---------

//in utils
func GetTypeObjects(ObjectsOc, ObjectsTypes []string) []string {
	// List of type of objects to export
	if len(ObjectsTypes) == 0 {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job"}
	} else if ObjectsTypes[0] == "default" {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job"}
	} else if ObjectsOc[0] == "all" {
		ObjectsTypes = []string{"service", "buildconfig", "build", "configmap", "daemonset","daemonset",
			"deployment", "deploymentconfig", "event","endpoints","horizontalpodautoscaler","imagestream",
			"imagestreamtag","ingress","group","job", "limitrange","node","namespace","pod","persistentvolume",
			"persistentvolumeclaim","policy","project","quota", "resourcequota","replicaset",
			"replicationcontroller","rolebinding","route","secret","serviceaccount","service","user"}
	} else {
		if len(ObjectsTypes) == 1 {
			ObjectsTypes = strings.Split(ObjectsTypes[0], ",")
		}
	}
	return ObjectsTypes

}

//in utils
func GetValueFromConfig(s string) interface{} {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName(".volume2volume")

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		value := viper.Get(s)
		if value != nil {
			return value
		}
	}
	return ""
}

//in utils
func GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
	UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		//fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if PathTemplate == ""{
				PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			// fmt.Println("PathTemplate -> " + PathTemplate)
		case "pathdata":
			if PathData == ""{
				PathData = GetValueFromConfig("PathData").(string)
			}
		case "clusterto":
			if ClusterTo == ""{
				ClusterTo = GetValueFromConfig("ClusterTo").(string)
			}
			//ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			if ClusterFrom == ""{
				ClusterFrom = GetValueFromConfig("ClusterFrom").(string)
			}
			//ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			if ProjectTo == ""{
				ProjectTo = GetValueFromConfig("ProjectTo").(string)
			}
			//ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			if ProjectFrom == ""{
				ProjectFrom = GetValueFromConfig("ProjectFrom").(string)
			}
		case "usernamefrom":
			if UsernameFrom == ""{
				UsernameFrom = GetValueFromConfig("UsernameFrom").(string)
			}
		case "usernameto":
			if UsernameTo == ""{
				UsernameTo = GetValueFromConfig("UsernameTo").(string)
			}
		case "passwordfrom":
			if PasswordFrom == ""{
				PasswordFrom = GetValueFromConfig("PasswordFrom").(string)
			}
		case "passwordto":
			if PasswordTo == ""{
				PasswordTo = GetValueFromConfig("PasswordTo").(string)
			}
		case "objects":
			if ObjectsOc[0] == "" {
				ObjectsOc = []string{GetValueFromConfig("objects").(string)}
				ObjectsOc = GetTypeObjects(ObjectsOc, ObjectsOc)
				//fmt.Println(ObjectsOc)
			}
		}
	}
}

//in utils
func GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
	UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) (string, string, string, string,
		string, string, string, string, string, string, []string) {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		// fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if PathTemplate == ""{
				PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			// fmt.Println("PathTemplate -> " + PathTemplate)
		case "pathdata":
			if PathData == ""{
				PathData = GetValueFromConfig("PathData").(string)
			}
		case "clusterto":
			if ClusterTo == ""{
				ClusterTo = GetValueFromConfig("ClusterTo").(string)
			}
			//ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			if ClusterFrom == ""{
				ClusterFrom = GetValueFromConfig("ClusterFrom").(string)
			}
			//ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			if ProjectTo == ""{
				ProjectTo = GetValueFromConfig("ProjectTo").(string)
			}
			//ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			if ProjectFrom == ""{
				ProjectFrom = GetValueFromConfig("ProjectFrom").(string)
			}
		case "usernamefrom":
			if UsernameFrom == ""{
				UsernameFrom = GetValueFromConfig("UsernameFrom").(string)
			}
		case "usernameto":
			if UsernameTo == ""{
				UsernameTo = GetValueFromConfig("UsernameTo").(string)
			}
		case "passwordfrom":
			if PasswordFrom == ""{
				PasswordFrom = GetValueFromConfig("PasswordFrom").(string)
			}
		case "passwordto":
			if PasswordTo == ""{
				PasswordTo = GetValueFromConfig("PasswordTo").(string)
			}
		case "objects":
			if ObjectsOc[0] == "" {
				ObjectsOc = []string{GetValueFromConfig("objects").(string)}
				ObjectsOc = GetTypeObjects(ObjectsOc, ObjectsOc)
				// fmt.Println(ObjectsOc)
			}
		}
	}
	return PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc
}

//in utils
func GetDeploymentReplicaSet(pod string) (string, string) {
	auxString := strings.Split(pod, "-")
	deploymentName := auxString[0]
	replicaSetName := deploymentName + "-" + auxString[1]
	return deploymentName, replicaSetName
}


// in utils
func createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName string,
	descriptionVolume, descriptionVolumeMount map[string]interface{}) map[string]interface{} {

	var m map[string]interface{}
	m = make(map[string]interface{})
	m["pathVolume"] = pathVolume
	m["volumeName"] = volumeName
	m["podName"] = podName
	m["mountPath"] = mountPath
	m["rsName"] = rsName
	m["deploymentName"] = deploymentName
	m["descriptionVolume"] = descriptionVolume
	m["descriptionVolumeMount"] = descriptionVolumeMount

	f, err3 := os.Create(pathVolume + "/data.json")

	if err3 != nil {
		fmt.Println("Error creating data.json")
		fmt.Println(err3)
	} else {
		objectOs, err2 := json.Marshal(m)
		if err2 != nil {
			fmt.Println("Error creating the json object")
			fmt.Println(err2)
		} else {
			f.WriteString(string(objectOs))
			f.Sync()
			fmt.Println("Created  data.json in " + pathVolume)
		}
	}
	return m
}

//in utils
func loginCluster(cluster, username, password string) {
	username = "--username=" + username
	password = "--password=" + password
	CmdLogin := exec.Command("oc", "login", cluster, username, password)
	//CmdLogin := exec.Command("oc", "login", cluster, "-u", "system:admin")
	CmdOut, err := CmdLogin.Output()
	fmt.Println(string(CmdOut))
	checkErrorMessage(err, "Error running login")
	fmt.Println(string(CmdOut))
}


//in utils
func changeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	checkErrorMessage(err, "Error running: change project")
	fmt.Println(string(CmdProjectOut))
}


//in utils
func checkErrorMessage(err error, message string){
	if err != nil {

		fmt.Println(message)
		//panic(err)
		fmt.Println(err)
		fmt.Println("Error managed by checkErrorMessage")
	}
}

// in utils
func getObjects(typeObject string) string {
	CmdGetDeployments := exec.Command("oc", "get", typeObject, "-o", "json")
	CmdOut, err := CmdGetDeployments.Output()
	if err != nil {
		fmt.Println("getObjects error in type " + typeObject)
		return ""
	}
	//checkErrorMessage(err, "Error running get " + typeObject)
	return string(CmdOut)
}







//END