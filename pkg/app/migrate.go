package app

import (
	"volume2volume/pkg/utils"
	"time"
	"fmt"
)

func MigrateVolume(deploymentName, volumeName string) {

	//Create Restic To
	auxPath := "./volumes/pairs/" + deploymentName + "/" + volumeName + "/"

	utils.CreateObject(auxPath + "resticTo.json")
	// TODO Wait a minut
	fmt.Println("sleeping")
	time.Sleep(1*time.Minute)
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


	//Create Recovery To
	utils.CreateObject(auxPath + "recoveryTo.json")

}

func Migrate() {
	// TODO

	// Take the full list

	// MigrateVolume
	// IDEA Use concurrency


}


