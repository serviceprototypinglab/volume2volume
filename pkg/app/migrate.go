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
	auxPath := PathData + "/pairs/" + deploymentName + "/" + volumeName + "/"

	utils.LoginCluster(ClusterTo, UsernameTo, PasswordTo)
	utils.ChangeProject(ProjectTo)
	utils.CreateObject(auxPath + "resticTo.json")
	// Check resticTo is done.
	fmt.Println("sleeping")
	time.Sleep(20*time.Second)
	fmt.Println("wake up")

	//Check if it is done
	// Deployment available
	// Restic re

	//Create Restic From
	utils.LoginCluster(ClusterFrom, UsernameFrom, PasswordFrom)
	utils.ChangeProject(ProjectFrom)
	utils.CreateObject(auxPath + "resticFrom.json")

	//Check if it is done
	// TODO
	// Status.BackupCount > 1
	fmt.Println("sleeping")
	time.Sleep(20*time.Second)
	fmt.Println("wake up")

	//Create Recovery To
	utils.LoginCluster(ClusterTo, UsernameTo, PasswordTo)
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


