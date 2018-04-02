package app


import (
	"fmt"
	"github.com/spf13/cobra"
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


func PairsVolumesByName(PathData, PathTemplate, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) ([]map[string]interface{}, []map[string]interface{})  {
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

func GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if PathTemplate == ""{
				PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			fmt.Println("PathTemplate -> " + PathTemplate)
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
				fmt.Println(ObjectsOc)
			}
		}
	}
}
func GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) (string, string, string, string, string, string, string, string, string, string, []string) {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if PathTemplate == ""{
				PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			fmt.Println("PathTemplate -> " + PathTemplate)
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
				fmt.Println(ObjectsOc)
			}
		}
	}
	return PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc
}


func GetDeploymentReplicaSet(pod string) (string, string) {
	auxString := strings.Split(pod, "-")
	deploymentName := auxString[0]
	replicaSetName := deploymentName + "-" + auxString[1]
	return deploymentName, replicaSetName
}

func ExportDataFromVolume(pod string, path string, mountPath string) {
	a := "oc rsync " + pod + ":" + mountPath + "/" +  " " + path + "/data"
	fmt.Println(a)
	cmdExportData := exec.Command("oc", "rsync", pod + ":" + mountPath + "/", path + "/data")
	cmdExportOut, err := cmdExportData.Output()
	if err != nil {
		fmt.Println("Error migrating " + a)
		fmt.Println(err)
	} else {
		fmt.Println(string(cmdExportOut))
	}
}

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

func changeProject(projectName string) {
	CmdProject := exec.Command("oc", "project", projectName)
	CmdProjectOut, err := CmdProject.Output()
	checkErrorMessage(err, "Error running: change project")
	fmt.Println(string(CmdProjectOut))
}

func checkErrorMessage(err error, message string){
	if err != nil {

		fmt.Println(message)
		//panic(err)
		fmt.Println(err)
		fmt.Println("Error managed by checkErrorMessage")
	}
}

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

func ExportData(cmd *cobra.Command, args []string, PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) {


	GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
	loginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	os.Mkdir(PathData, os.FileMode(0777)) //All permision??
	changeProject(ProjectFrom)

	var dat map[string]interface{}
	typeObject := "pods"
	typeString := getObjects(typeObject)
	byt := []byte(typeString)
	if err1 := json.Unmarshal(byt, &dat); err1 != nil {
		fmt.Println("Error with the objects with type " + typeObject)
		fmt.Println("-------")
		if typeString != "" {
			fmt.Println(typeString)
		}
	} else {
		items := dat["items"].([]interface{})
		if len(items) > 0 {
			os.Mkdir(PathData, os.FileMode(0777))

			var a [] map[string]interface{}

			//Take the name of the object
			for i := range items {
				var podName string
				nameObjectOsAux, ok :=
					items[i].(map[string]interface{})["metadata"].(map[string]interface{})["name"].(string)
				if ok {
					podName = nameObjectOsAux
				} else {
					podName = typeObject + string(i)

				}
				//Create a folder for each deployment
				deploymentName, rsName := GetDeploymentReplicaSet(podName)
				os.Mkdir(PathData+"/"+deploymentName, os.FileMode(0777))
				os.Mkdir(PathData+"/"+deploymentName+"/"+podName, os.FileMode(0777))
				//fmt.Println(podName)
				var volumeName string
				volumesAux, ok :=
					items[i].(map[string]interface{})["spec"].(map[string]interface{})["volumes"].([]interface{})
				if ok {
					for j := range volumesAux {
						volumeName = volumesAux[j].(map[string]interface{})["name"].(string)
						//fmt.Println(volumeName)
						descriptionVolume := volumesAux[j].(map[string]interface{})
						//fmt.Println(descriptionVolume)
						volumesMountAuxs, ok1 := items[i].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})
						for u := range volumesMountAuxs {
							if ok1 {
								volumesMountAux := volumesMountAuxs[u].(map[string]interface{})["volumeMounts"].([]interface{})
								for k := range volumesMountAux {
									nameVolumeMount := volumesMountAux[k].(map[string]interface{})["name"].(string)
									if nameVolumeMount == volumeName {
										descriptionVolumeMount := volumesMountAux[k].(map[string]interface{})
										mountPath := volumesMountAux[k].(map[string]interface{})["mountPath"].(string)
										pathVolume := PathData+"/"+deploymentName+"/"+podName + "/" + volumeName
										os.Mkdir(pathVolume, os.FileMode(0777))
										aux := createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
											descriptionVolume, descriptionVolumeMount)
										a = append(a, aux)
										os.Mkdir(pathVolume + "/data", os.FileMode(0777))
										//ExportDataFromVolume(podName, pathVolume, mountPath)
									}
								}
							}
						}
					}
				}
			}
			f, err3 := os.Create(PathData +"/data.json")
			if err3 != nil {
				fmt.Println("Error creating data.json")
				fmt.Println(err3)
			} else {
				objectOs, err2 := json.Marshal(a)
				if err2 != nil {
					fmt.Println("Error creating the json object")

					fmt.Println(err2)
				} else {
					f.WriteString(string(objectOs))
					f.Sync()
					fmt.Println("Created  data.json in" + PathData )
				}
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}


//"----------- UP ---"
func createRecovery(){

	// find all pairs of volumes and Resict objects

	// create the proper recovery object


}

func UpData(cmd *cobra.Command, args []string, PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) {

	GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
	loginCluster(ClusterTo, UsernameTo, PasswordTo)
	os.Mkdir(PathData, os.FileMode(0777)) //All permission??
	changeProject(ProjectTo)


	data := ReadJsonData("./volumes")

	var dat map[string]interface{}
	typeObject := "pods"
	typeString := getObjects(typeObject)
	byt := []byte(typeString)
	if err1 := json.Unmarshal(byt, &dat); err1 != nil {
		fmt.Println("Error with the objects with type " + typeObject)
		fmt.Println("-------")
		if typeString != "" {
			fmt.Println(typeString)
		}
	} else {
		items := dat["items"].([]interface{})
		if len(items) > 0 {
			//Take the name of the object
			for i := range items {
				var podName string
				nameObjectOsAux, ok :=
					items[i].(map[string]interface{})["metadata"].
					(map[string]interface{})["name"].(string)
				if ok {
					podName = nameObjectOsAux
				} else {
					podName = typeObject + string(i)

				}
				//Create a folder for each deployment
				deploymentName, _ := GetDeploymentReplicaSet(podName)

				//FIND DEPLOYMENT AND PROJECT NAME
				for _, a := range data {
					if a["deploymentName"] == deploymentName {
						path := "./volumes/" + deploymentName + "/" + a["podName"].(string) + "/" +
							a["volumeName"].(string)
						mountPath := a["mountPath"].(string)
						UpDataToVolume(podName, path, mountPath)
					}
				}
				/*
					for _, v := range listDeployments() {
						deployment := getDeploymentName(v)
						if deployment == deploymentName {
							for _, podName := range listPods(deploymentName) {

								volumes := listVolumes(deploymentName+"/"+getPodNameFromPath(podName))
								for _, volumePath := range volumes {
									fmt.Println(volumePath)
								}
							}
						}
					}
					//fmt.Println(podName)
					/*var volumeName string
					volumesAux, ok :=
						items[i].(map[string]interface{})["spec"].(map[string]interface{})["volumes"].([]interface{})
					if ok {
						for j := range volumesAux {
							volumeName = volumesAux[j].(map[string]interface{})["name"].(string)
							//fmt.Println(volumeName)
							descriptionVolume := volumesAux[j].(map[string]interface{})
							//fmt.Println(descriptionVolume)
							volumesMountAuxs, ok1 := items[i].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})
							for u := range volumesMountAuxs {
								if ok1 {
									volumesMountAux := volumesMountAuxs[u].(map[string]interface{})["volumeMounts"].([]interface{})
									for k := range volumesMountAux {
										nameVolumeMount := volumesMountAux[k].(map[string]interface{})["name"].(string)
										if nameVolumeMount == volumeName {
											descriptionVolumeMount := volumesMountAux[k].(map[string]interface{})
											mountPath := volumesMountAux[k].(map[string]interface{})["mountPath"].(string)
											pathVolume := PathData+"/"+deploymentName+"/"+podName + "/" + volumeName
											os.Mkdir(pathVolume, os.FileMode(0777))
											createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
												descriptionVolume, descriptionVolumeMount)
											os.Mkdir(pathVolume + "/data", os.FileMode(0777))
											exportDataFromVolume(podName, pathVolume, mountPath)
										}
									}
								}
							}
						}
					}*/
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}

func UpDataToVolume(podName, path, mountPath string) {
	a := "oc rsync " +  path + "/data/"  +  " " + podName + ":" + mountPath + "/"
	fmt.Println(a)
	cmdUpData := exec.Command("oc", "rsync", path + "/data/", podName + ":" + mountPath + "/")
	cmdUpOut, err := cmdUpData.Output()
	if err != nil {
		fmt.Println("Error migrating " + a)
		fmt.Println(err)
	} else {
		fmt.Println(string(cmdUpOut))
	}
}

//END