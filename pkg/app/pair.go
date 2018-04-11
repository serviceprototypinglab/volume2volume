package app

import (
	"fmt"
	"volume2volume/pkg/utils"
	"os"
	"os/exec"
)

func PairsVolumesByName(PathData, PathTemplate, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo,
UsernameFrom, PasswordFrom, PasswordTo string, ObjectsOc []string) ([]map[string]interface{},
	[]map[string]interface{}) {
	//Read  Volumes/ClusterFrom/data.json
	if PathData == "" {
		PathData = "volumes"
	}
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
					fmt.Println("PAIR!!!!!")
					//fmt.Println(v["volumeName"])
					from = append(from, v)
					to = append(to, k)
					pathFrom := v["pathVolume"].(string)
					pathTo := k["pathVolume"].(string)
					deploymentName := v["deploymentName"].(string)
					os.Mkdir(PathData + "/pairs", os.FileMode(0777))
					os.Mkdir(PathData + "/pairs/" + deploymentName, os.FileMode(0777))
					copy(pathFrom , PathData + "/pairs/" + deploymentName)
					copy(pathTo , PathData + "/pairs/" + deploymentName)
				}
			}
		}
	}
	return from, to
}




func copy(srcFolder string, destFolder string) {
	// Read all content of src to data
	/*data, err := ioutil.ReadFile(src)
	fmt.Println(err)
	fmt.Println("readddd")
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	fmt.Println(err)
	fmt.Println("copied")*/

	cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
	err := cpCmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

