package app

import (
	"volume2volume/pkg/utils"
)

func BackUpVolume(PathData, deploymentName, volumeName string) {

	//Create Restic To
	auxPath := PathData + "/pairs/" + deploymentName + "/" + volumeName + "/"
	//Create Restic From
	utils.CreateObject(auxPath + "resticFrom.json")
}

func BackUp(PathData string) {
	var pairs []map[string]interface{}
	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	for _, v := range pairs {
		PrintVolumes(v)
		// TEST THAT
		go BackUpVolume(PathData, v["deploymentName"].(string), v["volumeName"].(string))
	}
}


