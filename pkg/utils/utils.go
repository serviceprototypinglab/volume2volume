/*
	In this package you can find useful methods which interacting with the openshift cluster.
*/
package utils

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"volume2volume/pkg/confObject"
	"strings"
	"reflect"
)

// GET CONFIGURATION
func GetTypeObjects(ObjectsOc, ObjectsTypes []string) []string {
	// List of type of objects to export
	if len(ObjectsTypes) == 0 {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job"}
	} else if ObjectsTypes[0] == "default" {
		ObjectsTypes = []string{"service", "deployment", "secrets", "configmap", "job"}
	} else if ObjectsOc[0] == "all" {
		ObjectsTypes = []string{"service", "buildconfig", "build", "configmap", "daemonset", "daemonset",
			"deployment", "deploymentconfig", "event", "endpoints", "horizontalpodautoscaler", "imagestream",
			"imagestreamtag", "ingress", "group", "job", "limitrange", "node", "namespace", "pod", "persistentvolume",
			"persistentvolumeclaim", "policy", "project", "quota", "resourcequota", "replicaset",
			"replicationcontroller", "rolebinding", "route", "secret", "serviceaccount", "service", "user"}
	} else {
		if len(ObjectsTypes) == 1 {
			ObjectsTypes = strings.Split(ObjectsTypes[0], ",")
		}
	}
	return ObjectsTypes
}

// Get the configuration using the key: "s"  from the config file ~/.volume2volume.yaml
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

func GetAllValueFromObject(confObject confObject.ConfObject) {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		//fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if confObject.PathTemplate == ""{
				confObject.PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			// fmt.Println("PathTemplate -> " + PathTemplate)
		case "pathdata":
			if confObject.PathData == ""{
				confObject.PathData = GetValueFromConfig("PathData").(string)
			}
		case "clusterto":
			if confObject.ClusterTo == ""{
				confObject.ClusterTo = GetValueFromConfig("ClusterTo").(string)
			}
			//ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			if confObject.ClusterFrom == ""{
				confObject.ClusterFrom = GetValueFromConfig("ClusterFrom").(string)
			}
			//ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			if confObject.ProjectTo == ""{
				confObject.ProjectTo = GetValueFromConfig("ProjectTo").(string)
			}
			//ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			if confObject.ProjectFrom == ""{
				confObject.ProjectFrom = GetValueFromConfig("ProjectFrom").(string)
			}
		case "usernamefrom":
			if confObject.UsernameFrom == ""{
				confObject.UsernameFrom = GetValueFromConfig("UsernameFrom").(string)
			}
		case "usernameto":
			if confObject.UsernameTo == ""{
				confObject.UsernameTo = GetValueFromConfig("UsernameTo").(string)
			}
		case "passwordfrom":
			if confObject.PasswordFrom == ""{
				confObject.PasswordFrom = GetValueFromConfig("PasswordFrom").(string)
			}
		case "passwordto":
			if confObject.PasswordTo == ""{
				confObject.PasswordTo = GetValueFromConfig("PasswordTo").(string)
			}
		case "objects":
			if confObject.ObjectsOc[0] == "" {
				confObject.ObjectsOc = []string{GetValueFromConfig("objects").(string)}
				confObject.ObjectsOc = GetTypeObjects(confObject.ObjectsOc, confObject.ObjectsOc)
				//fmt.Println(ObjectsOc)
			}
		}
	}
}

func GetAllValueReturnObject(confObject confObject.ConfObject) confObject.ConfObject {
	keys := []string{"pathtemplate","pathdata","objects","clusterto", "clusterfrom","projectto", "projectfrom",
		"usernamefrom", "usernameto", "passwordfrom", "passwordto"}
	for _, keyConfig := range keys {
		//fmt.Println("-------")
		//fmt.Println(keyConfig)
		//fmt.Println(viper.GetString(keyConfig))
		//fmt.Println("key -> " + keyConfig)
		switch keyConfig {
		case "pathtemplate":
			if confObject.PathTemplate == ""{
				confObject.PathTemplate = GetValueFromConfig("PathTemplate").(string)
			}
			// fmt.Println("PathTemplate -> " + PathTemplate)
		case "pathdata":
			if confObject.PathData == ""{
				confObject.PathData = GetValueFromConfig("PathData").(string)
			}
		case "clusterto":
			if confObject.ClusterTo == ""{
				confObject.ClusterTo = GetValueFromConfig("ClusterTo").(string)
			}
			//ClusterTo = viper.GetString(keyConfig)
		case "clusterfrom":
			if confObject.ClusterFrom == ""{
				confObject.ClusterFrom = GetValueFromConfig("ClusterFrom").(string)
			}
			//ClusterFrom = viper.GetString(keyConfig)
		case "projectto":
			if confObject.ProjectTo == ""{
				confObject.ProjectTo = GetValueFromConfig("ProjectTo").(string)
			}
			//ProjectTo = viper.GetString(keyConfig)
		case "projectfrom":
			if confObject.ProjectFrom == ""{
				confObject.ProjectFrom = GetValueFromConfig("ProjectFrom").(string)
			}
		case "usernamefrom":
			if confObject.UsernameFrom == ""{
				confObject.UsernameFrom = GetValueFromConfig("UsernameFrom").(string)
			}
		case "usernameto":
			if confObject.UsernameTo == ""{
				confObject.UsernameTo = GetValueFromConfig("UsernameTo").(string)
			}
		case "passwordfrom":
			if confObject.PasswordFrom == ""{
				confObject.PasswordFrom = GetValueFromConfig("PasswordFrom").(string)
			}
		case "passwordto":
			if confObject.PasswordTo == ""{
				confObject.PasswordTo = GetValueFromConfig("PasswordTo").(string)
			}
		case "objects":
			if confObject.ObjectsOc[0] == "" {
				confObject.ObjectsOc = []string{GetValueFromConfig("objects").(string)}
				confObject.ObjectsOc = GetTypeObjects(confObject.ObjectsOc, confObject.ObjectsOc)
				//fmt.Println(ObjectsOc)
			}
		}
	}
	return confObject
}

// Get ALL the configuration using the key: "s"  from the config file ~/.volume2volume.yaml
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

// Get ALL the configuration using the key: "s"  from the config file ~/.volume2volume.yaml
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

func GetDeploymentReplicaSet(pod string) (string, string) {
	auxString := strings.Split(pod, "-")
	deploymentName := auxString[0]
	replicaSetName := deploymentName + "-" + auxString[1]
	return deploymentName, replicaSetName
}


func In_array(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

