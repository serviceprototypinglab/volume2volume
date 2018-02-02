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
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"encoding/json"
	"os/exec"
	"strings"
)

var (
	ClusterFrom  string
	ClusterTo    string
	ProjectFrom  string
	ProjectTo    string
	PathTemplate string
	PathData     string
	UsernameFrom string
	UsernameTo   string
	PasswordFrom string
	PasswordTo   string
	cfgFile      string
)
var ObjectsOc []string
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "volume2volume",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.volume2volume.yaml)")


	RootCmd.PersistentFlags().StringVarP(&ClusterFrom, "clusterFrom", "", "", "Cluster where is the project that you want to migrate")
	RootCmd.PersistentFlags().StringVarP(&ClusterTo, "clusterTo", "", "", "Cluster where you want to migrate the project")
	RootCmd.PersistentFlags().StringVarP(&ProjectFrom, "projectFrom", "", "", "name of the old Openshift project")
	RootCmd.PersistentFlags().StringVarP(&ProjectTo, "projectTo", "", "", "name of the new Openshift project")
	RootCmd.PersistentFlags().StringVarP(&UsernameFrom, "usernameFrom", "", "", "username in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&UsernameTo, "usernameTo", "", "", "username in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PasswordFrom, "passwordFrom", "", "", "password in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&PasswordTo, "passwordTo", "", "", "password in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PathTemplate, "pathTemplate","","", "path where export the templates")
	RootCmd.PersistentFlags().StringVarP(&PathData, "pathData","", "", "path where export the volumes")
	defaultValue := []string{""}
	RootCmd.PersistentFlags().StringArrayVarP(&ObjectsOc, "objects", "o", defaultValue, "list of objects to export" )

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".volume2volume" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".volume2volume")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getTypeObjects(ObjectsTypes []string) []string {
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

func getValueFromConfig(s string) interface{} {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	viper.AddConfigPath(home)
	viper.SetConfigName(".os2os")

	if err := viper.ReadInConfig(); err == nil {
		//fmt.Println("Using config file:", viper.ConfigFileUsed())
		value := viper.Get(s)
		if value != nil {
			return value
		}
	}
	return ""
}

func getAllValue(){
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))

		switch keyConfig {
		case "pathtemplate":
			if PathTemplate == ""{
				PathTemplate = getValueFromConfig("PathTemplate").(string)
			}
		case "pathdata":
			if PathData == ""{
				PathData = getValueFromConfig("PathData").(string)
			}
		case "clusterto":
			if ClusterTo == ""{
				ClusterTo = getValueFromConfig("ClusterTo").(string)
			}
			//ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			if ClusterFrom == ""{
				ClusterFrom = getValueFromConfig("ClusterFrom").(string)
			}
			//ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			if ProjectTo == ""{
				ProjectTo = getValueFromConfig("ProjectTo").(string)
			}
			//ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			if ProjectFrom == ""{
				ProjectFrom = getValueFromConfig("ProjectFrom").(string)
			}
		case "usernamefrom":
			if UsernameFrom == ""{
				UsernameFrom = getValueFromConfig("UsernameFrom").(string)
			}
		case "usernameto":
			if UsernameTo == ""{
				UsernameTo = getValueFromConfig("UsernameTo").(string)
			}
		case "passwordfrom":
			if PasswordFrom == ""{
				PasswordFrom = getValueFromConfig("PasswordFrom").(string)
			}
		case "passwordto":
			if PasswordTo == ""{
				PasswordTo = getValueFromConfig("PasswordTo").(string)
			}
		case "objects":
			if ObjectsOc[0] == "" {
				ObjectsOc = []string{getValueFromConfig("objects").(string)}
				ObjectsOc = getTypeObjects(ObjectsOc)
				fmt.Println(ObjectsOc)
			}
		}
	}
}


func getDeploymentReplicaSet(pod string) (string, string) {
	auxString := strings.Split(pod, "-")
	deploymentName := auxString[0]
	replicaSetName := deploymentName + "-" + auxString[1]
	return deploymentName, replicaSetName
}

func exportDataFromVolume(pod string, path string, mountPath string) {
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

func exportData(cmd *cobra.Command, args []string) {


	getAllValue()
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
				deploymentName, rsName := getDeploymentReplicaSet(podName)
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
										//exportDataFromVolume(podName, pathVolume, mountPath)
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


