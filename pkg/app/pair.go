package app

import (
	"fmt"
	"volume2volume/pkg/utils"
)

func PairsVolumesByName(PathData, PathTemplate, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo,
UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) ([]map[string]interface{},
	[]map[string]interface{}) {
	//Read  Volumes/ClusterFrom/data.json
	var from [] map[string]interface{}
	var to [] map[string]interface{}
	utils.GetAllValue(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
		UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
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

func Pair(from, to [] map[string]interface{}) {
	for i, v := range from {
		fmt.Println("volume -> %d", i)
		fmt.Println(v)
		fmt.Println(to[i])
	}
}