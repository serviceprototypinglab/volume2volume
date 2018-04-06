package app

import (
	"os"
	"fmt"
	"encoding/json"
	"volume2volume/pkg/utils"

)

func FindVolumes(cluster, PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
	UsernameTo, UsernameFrom, PasswordFrom, PasswordTo  string,ObjectsOc []string) {
	GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
	var cluster1 string
	var project1 string
	if cluster == "ClusterFrom"{
		cluster1 = ClusterFrom
		project1 = ProjectFrom
	} else {
		cluster1 = ClusterTo
		project1 = ProjectTo
	}

	fmt.Println("USER -> " +  UsernameFrom)
	loginCluster(cluster1, UsernameFrom, PasswordFrom)
	os.Mkdir(PathData, os.FileMode(0777)) //All permission?
	os.Mkdir(PathData + "/" + cluster, os.FileMode(0777))

	changeProject(project1)

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
			os.Mkdir(PathData+ "/" + cluster, os.FileMode(0777))

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
				os.Mkdir(PathData+ "/" + cluster +"/"+deploymentName, os.FileMode(0777))
				os.Mkdir(PathData+ "/" + cluster +"/"+deploymentName+"/"+podName, os.FileMode(0777))
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
										pathVolume := PathData+ "/" + cluster +"/"+deploymentName+"/"+podName + "/" + volumeName
										os.Mkdir(pathVolume, os.FileMode(0777))
										aux := createJson(pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
											descriptionVolume, descriptionVolumeMount)
										a = append(a, aux)
										pathRestic := pathVolume + "/restic"
										os.Mkdir(pathRestic, os.FileMode(0777))
										//TODO create restic
										createRestic(ProjectFrom, volumeName, deploymentName, mountPath, pathRestic)
										//ExportDataFromVolume(podName, pathVolume, mountPath)
									}
								}
							}
						}
					}
				}
			}
			f, err3 := os.Create(PathData + "/" + cluster +"/data.json")
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
					fmt.Println("Created  data.json in " + PathData + "/" + cluster  )
				}
			}
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}



func createRestic(namespace, volumeName, deploymentName, mountPath, pathRestic string) {
	/*type restic struct {
ReadJsonData
	}*/

	restic := utils.ReadJson("templates", "restic_template")
	fmt.Println(restic)
	//Change name
	auxName := "restic-" + deploymentName
	fmt.Println("name -->")
	restic["metadata"].(map[string]interface{})["name"] = auxName
	//fmt.Println(restic["metadata"].(map[string]interface{})["name"].(string))
	//Change namespace
	restic["metadata"].(map[string]interface{})["namespace"] = namespace
	//Change volumeName
	restic["spec"].(map[string]interface{})["volumeMounts"].([]interface{})[0].(map[string]interface{})["name"] = volumeName
	//Change deploymentName
	restic["spec"].(map[string]interface{})["volumeMounts"].([]interface{})[0].(map[string]interface{})["mountPath"] = mountPath
	//change mountPath
	restic["spec"].(map[string]interface{})["fileGroups"].([]interface{})[0].(map[string]interface{})["path"] = mountPath

	fmt.Println(restic)
    // TODO
    // Backend -> local, s3, glusterFS, ...

	err := utils.WriteJson(pathRestic, "restic", restic)
	if err != nil {
		fmt.Println("Error creating " + auxName)
	}
    //write json in path restic
	/*f, err3 := os.Create(pathRestic +"/restic.json")
	if err3 != nil {
		fmt.Println("Error creating data.json")
		fmt.Println(err3)
	} else {
		objectOs, err2 := json.Marshal(restic)
		if err2 != nil {
			fmt.Println("Error creating the json object")

			fmt.Println(err2)
		} else {
			f.WriteString(string(objectOs))
			f.Sync()
			fmt.Println("Created  data.json in" + pathRestic )
		}*/
	}





