package app

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
	"volume2volume/pkg/utils"
	"encoding/json"
	"os/exec"
)

//"----------- UP --------"
func createRecovery(){

	// find all pairs of volumes and Resict objects

	// create the proper recovery object


}

func UpData(cmd *cobra.Command, args []string, PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) {

	GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
	loginCluster(ClusterTo, UsernameTo, PasswordTo)
	os.Mkdir(PathData, os.FileMode(0777)) //All permission??
	changeProject(ProjectTo)


	data := utils.ReadJsonData("./volumes")

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
