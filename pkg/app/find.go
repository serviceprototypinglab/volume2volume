package app

import (
	"os"
	"fmt"
	"encoding/json"
	"volume2volume/pkg/utils"
	"volume2volume/pkg/confObject"
)


// Find all the volumes in the two clusters. Create a restic, a recovery a description and stats of each volume.
// Pairs the volumes by name in the same deployment (in cluster).
func FindVolumes(cluster, PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
	UsernameTo, UsernameFrom, PasswordFrom, PasswordTo  string,ObjectsOc []string) {

	// Get values
	utils.GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
		UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)

	// Choose for cluster From or cluster To
	var cluster1 string
	var project1 string
	if cluster == "ClusterFrom"{
		cluster1 = ClusterFrom
		project1 = ProjectFrom
		utils.LoginCluster(cluster1, UsernameFrom, PasswordFrom)

	} else {
		cluster1 = ClusterTo
		project1 = ProjectTo
		utils.LoginCluster(cluster1, UsernameTo, PasswordTo)

	}

	// Connect to the cluster
	fmt.Println("USER -> " +  UsernameFrom)
	// utils.LoginCluster(cluster1, UsernameFrom, PasswordFrom)
	os.Mkdir(PathData, os.FileMode(0777)) //All permission?
	os.Mkdir(PathData + "/" + cluster, os.FileMode(0777))

	// Go to the project
	utils.ChangeProject(project1)


	// Get pods
	var dat map[string]interface{}
	typeObject := "pods"
	typeString := utils.GetObjects(typeObject)
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

				var dataType string
				dataTypeAux, ok :=
					items[i].(map[string]interface{})["metadata"].(map[string]interface{})["labels"].(map[string]interface{})["v2v"].(string)
				if ok {
					fmt.Println("---------------->" + dataTypeAux)
					dataType = dataTypeAux
				} else {
					dataType = "Not Found"

				}
				//Create a folder for each deployment
				deploymentName, rsName := utils.GetDeploymentReplicaSet(podName)
				os.Mkdir(PathData+ "/" + cluster +"/"+deploymentName, os.FileMode(0777))
				os.Mkdir(PathData+ "/" + cluster +"/"+deploymentName+"/"+podName, os.FileMode(0777))
				//fmt.Println(podName)
				var volumeName string
				volumesAux, ok :=
					items[i].(map[string]interface{})["spec"].(map[string]interface{})["volumes"].([]interface{})
				if ok {
					// get volumes
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

										// Create description
										aux := CreateDescription(dataType, cluster, pathVolume, volumeName, podName, mountPath, rsName, deploymentName,
											descriptionVolume, descriptionVolumeMount)

										// Create recovery
										CreateRecovery(cluster, project1, volumeName, deploymentName, mountPath, pathVolume)
										// Create restic
										CreateRestic(cluster, project1, volumeName, deploymentName, mountPath, pathVolume)

										//ExportDataFromVolume(podName, pathVolume, mountPath)
										aux["size"] = CreateStats(cluster, project1, volumeName, deploymentName, mountPath, pathVolume, podName)
										a = append(a, aux)

									}
								}
							}
						}
					}
				}
			}


			// TODO CHANGE THAT ALSO
			err1 := utils.WriteJsonArray(PathData + "/" + cluster, "data", a)
			if err1 != nil {
				fmt.Println("Error creating " + "description")
			}

			/*f, err3 := os.Create(PathData + "/" + cluster +"/data.json")
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
			}*/
		} else {
			fmt.Println("No objects for the type " + typeObject)
		}
	}
}

// Create a json file with the description of the volume.
func CreateDescription(dataType, cluster, pathVolume, volumeName, podName, mountPath, rsName, deploymentName string,
	descriptionVolume, descriptionVolumeMount map[string]interface{}) map[string]interface{} {

	var nameJson string
	if cluster == "ClusterFrom"{
		nameJson = "descriptionFrom"
	} else {
		nameJson = "descriptionTo"
	}


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
	m["dataType"] = dataType


	/*err := utils.WriteJson(pathVolume, "data", m)
	if err != nil {
		fmt.Println("Error creating " + "data")
	}*/

	err1 := utils.WriteJson(pathVolume, nameJson, m)
	if err1 != nil {
		fmt.Println("Error creating " + "description")
	}

	/*f, err3 := os.Create(pathVolume + "/data.json")

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

	f1, err4 := os.Create(pathVolume + "/description.json")

	if err4 != nil {
		fmt.Println("Error creating data.json")
		fmt.Println(err4)
	} else {
		objectOs, err2 := json.Marshal(m)
		if err2 != nil {
			fmt.Println("Error creating the json object")
			fmt.Println(err2)
		} else {
			f1.WriteString(string(objectOs))
			f1.Sync()
			fmt.Println("Created  data.json in " + pathVolume)
		}
	}*/
	return m
}

// Create a json of the restic object for the volume given.
func CreateRestic(cluster, namespace, volumeName, deploymentName, mountPath, pathRestic string) {

	/*type restic struct {
ReadJsonData
	}*/


	var restic map[string]interface{}
	var nameRestic string
	// TODO Backend -> local, s3, glusterFS, ...
	if cluster == "ClusterFrom" {
		restic = utils.ReadJson("templates/restic", "restic_s3_template_from")
		nameRestic = "resticFrom"
	} else {
		restic = utils.ReadJson("templates/restic", "restic_s3_template_to")
		nameRestic = "resticTo"
	}

	fmt.Println(restic)
	//Change name
	auxName := "restic-" + deploymentName
	fmt.Println("name -->")
	auxName = deploymentName
	restic["metadata"].(map[string]interface{})["name"] = auxName
	//fmt.Println(restic["metadata"].(map[string]interface{})["name"].(string))
	//Change namespace
	restic["metadata"].(map[string]interface{})["namespace"] = namespace

	restic["spec"].(map[string]interface{})["selector"].(map[string]interface{})["matchLabels"].(map[string]interface{})["app"] = deploymentName
	//Change volumeName
	restic["spec"].(map[string]interface{})["volumeMounts"].([]interface{})[0].(map[string]interface{})["name"] = volumeName
	//Change deploymentName
	restic["spec"].(map[string]interface{})["volumeMounts"].([]interface{})[0].(map[string]interface{})["mountPath"] = mountPath
	//change mountPath
	restic["spec"].(map[string]interface{})["fileGroups"].([]interface{})[0].(map[string]interface{})["path"] = mountPath

	fmt.Println(restic)


	err := utils.WriteJson(pathRestic, nameRestic, restic)
	if err != nil {
		fmt.Println("Error creating " + auxName)
	}
}

// Create a json of the recovery object for the volume given.
func CreateRecovery(cluster, namespace, volumeName, deploymentName, mountPath, pathRestic string) {
	var recovery map[string]interface{}
	var nameRecovery string
	// TODO Backend -> local, s3, glusterFS, ...
	if cluster == "ClusterFrom" {
		recovery = utils.ReadJson("templates/recovery", "recovery_s3_template_from")
		nameRecovery= "recoveryFrom"
	} else {
		recovery = utils.ReadJson("templates/recovery", "recovery_s3_template_to")
		nameRecovery = "recoveryTo"
	}

	// Change namespace, name,
	auxName := "recovery-" + deploymentName
	auxName = deploymentName
	recovery["metadata"].(map[string]interface{})["name"] = auxName
	recovery["metadata"].(map[string]interface{})["namespace"] = namespace
	recovery["spec"].(map[string]interface{})["workload"].(map[string]interface{})["name"] = deploymentName
	recovery["spec"].(map[string]interface{})["paths"].([]interface{})[0] = mountPath
	recovery["spec"].(map[string]interface{})["recoveredVolumes"].([]interface{})[0].(map[string]interface{})["mountPath"] = mountPath

	err := utils.WriteJson(pathRestic, nameRecovery, recovery)
	if err != nil {
		fmt.Println("Error creating " + auxName)
	}
}

// Create a json with the stats of the volume.
func CreateStats(cluster, namespace, volumeName, deploymentName, mountPath, pathRestic, podName string) string {
	var stats map[string]interface{}
	var nameStats string
	if cluster == "ClusterFrom" {
		stats = utils.ReadJson("templates/stats", "stats_template_from")
		nameStats = "statsFrom"
	} else {
		stats = utils.ReadJson("templates/stats", "stats_template_to")
		nameStats = "statsTo"
	}

	auxName := "stats-" + deploymentName
	sizeVolume := utils.GetSizeVolume(podName, volumeName, mountPath)
	stats["name"]  = auxName
	stats["size"] = sizeVolume
	err := utils.WriteJson(pathRestic, nameStats, stats)
	if err != nil {
		fmt.Println("Error creating " + auxName)
	}
	return sizeVolume
}

func getItems() []interface{} {
	// TODO
	return nil
}

func getVolumes() {
	//TODO
}

func FindVolumes1(cluster string, conf confObject.ConfObject) error {
	// TODO
	return nil
}