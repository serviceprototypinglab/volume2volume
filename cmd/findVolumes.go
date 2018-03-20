// Copyright © 2018 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/spf13/cobra"
	"os"
	"io/ioutil"
)

// findVolumesCmd represents the findVolumes command
var findVolumesCmd = &cobra.Command{
	Use:   "findVolumes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("findVolumes called")
		//findAllVolumes(cmd, args)
		a,b := pairsVolumesByName()
		fmt.Println(a[0]["podName"])
		fmt.Println("---")
		fmt.Println(b[0]["podName"])
	},
}

func init() {
	RootCmd.AddCommand(findVolumesCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findVolumesCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findVolumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func findAllVolumes(cmd *cobra.Command, args []string){
	getAllValue()
	findVolumes("ClusterFrom")
	findVolumes("ClusterTo")
}

func findVolumes(cluster string) {
	getAllValue()
	var cluster1 string
	var project1 string
	if cluster == "ClusterFrom"{
		cluster1 = ClusterFrom
		project1 = ProjectFrom
	} else {
		cluster1 = ClusterTo
		project1 = ProjectTo
	}

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
				deploymentName, rsName := getDeploymentReplicaSet(podName)
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
										os.Mkdir(pathVolume + "/data", os.FileMode(0777))
										//exportDataFromVolume(podName, pathVolume, mountPath)
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

func pairsVolumesByName() ([]map[string]interface{}, []map[string]interface{})  {
	//Read  Volumes/ClusterFrom/data.json
	var from [] map[string]interface{}
	var to [] map[string]interface{}
	getAllValue()
	clusterFromVolumes := readJsonData(PathData + "/ClusterFrom")
	clusterToVolumes := readJsonData(PathData + "/ClusterTo")
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

func readJsonData(path string) []map[string]interface{} {
	fmt.Println(path)
	plan, _ := ioutil.ReadFile(path + "/data.json")
	//fmt.Println(plan)
	//var data []interface{}
	var data []map[string]interface{}
	err := json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error reading json")
		//fmt.Println(data)
		fmt.Println(err)
	}
	return data
}
