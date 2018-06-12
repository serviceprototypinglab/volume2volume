package app

import (
	"volume2volume/pkg/utils"
	"time"
	"fmt"
)

func MigrateVolume(PathData, deploymentName, volumeName,
	ClusterFrom, UsernameFrom, PasswordFrom, ProjectFrom,
		ClusterTo, UsernameTo, PasswordTo, ProjectTo string) {

	//Create Restic To
	auxPath := "./" + PathData + "/pairs/" + deploymentName + "/" + volumeName + "/"


	// TODO MAYBE CHANGE THAT
	utils.LoginAdmin(ClusterTo)
	utils.ChangeProject(ProjectTo)
	utils.CreateObject(auxPath + "resticTo.json")
	// Check resticTo is done.
	fmt.Println("RESTIC_TO")
	//time.Sleep(20*time.Second)
	//fmt.Println("wake up")

	//Check if it is done
	// Deployment available
	// Restic re

	//Create Restic From
	//utils.LoginAdmin(ClusterFrom)

	utils.ChangeProject(ProjectFrom)
	utils.CreateObject(auxPath + "resticFrom.json")
	fmt.Println("RESTIC_FROM")


	//Check if it is done
	// TODO
	// Status.BackupCount > 1
	//fmt.Println("sleeping")
	// kubectl describe restic "restic_name"

	time.Sleep(150*time.Second)
	//fmt.Println("wake up")

	//Create Recovery To
	//utils.LoginAdmin(ClusterTo)
	fmt.Println("Recovery_to")
	utils.ChangeProject(ProjectTo)
	utils.CreateObject(auxPath + "recoveryTo.json")

}

func Migrate(PathData, ClusterFrom, UsernameFrom, PasswordFrom, ProjectFrom, ClusterTo,
	UsernameTo, PasswordTo, ProjectTo string) {
	var pairs []map[string]interface{}
	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	for _, v := range pairs {
		PrintVolumes(v)
		// TEST THAT
		MigrateVolume(PathData, v["deploymentName"].(string), v["volumeName"].(string),
			ClusterFrom, UsernameFrom, PasswordFrom, ProjectFrom,
				ClusterTo, UsernameTo, PasswordTo, ProjectTo)
	}
}


