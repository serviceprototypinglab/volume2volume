package app

import (
	"volume2volume/pkg/utils"
)


// Create the object restic for the volume: VolumeName in the deployment: deploymentName.
func BackUpVolume(PathData, deploymentName, volumeName string) {

	//Create Restic To
	auxPath := PathData + "/pairs/" + deploymentName + "/" + volumeName + "/"
	//Create Restic From
	utils.CreateObject(auxPath + "resticFrom.json")
}

// Create all the objects restic for all the pairs of volumes.
func BackUp(PathData string) {
	var pairs []map[string]interface{}
	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	for _, v := range pairs {
		PrintVolumes(v)
		// TEST THAT
		go BackUpVolume(PathData, v["deploymentName"].(string), v["volumeName"].(string))
	}
}


