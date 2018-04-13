package app

import (
	"volume2volume/pkg/utils"
	"time"
	"fmt"
)

func MigrateVolume(PathData, deploymentName, volumeName string) {

	//Create Restic To
	auxPath := PathData + "/pairs/" + deploymentName + "/" + volumeName + "/"

	utils.CreateObject(auxPath + "resticTo.json")
	// TODO Wait a minut
	fmt.Println("sleeping")
	time.Sleep(20*time.Second)
	fmt.Println("wake up")

	//Check if it is done
	// TODO
	// Deployment available
	// Restic re

	//Create Restic From
	utils.CreateObject(auxPath + "resticFrom.json")

	//Check if it is done
	// TODO
	// Status.BackupCount > 1
	fmt.Println("sleeping")
	time.Sleep(20*time.Second)
	fmt.Println("wake up")

	//Create Recovery To
	utils.CreateObject(auxPath + "recoveryTo.json")

}

func Migrate(PathData string) {
	var pairs []map[string]interface{}
	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	for _, v := range pairs {
		PrintVolumes(v)
		// TEST THAT
		go MigrateVolume(PathData, v["deploymentName"].(string), v["volumeName"].(string))
	}
}


