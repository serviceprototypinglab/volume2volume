package app

import (
	"fmt"
	"volume2volume/pkg/utils"
)

// Show the volumes found it and the stats.
func ShowVolumes(PathData string) {
	var pairs []map[string]interface{}
	//TODO Show all the volumes pairs and statistics in a nice way.

	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")

	fmt.Println("PAIRS OF VOLUMES")
	fmt.Println("-------*------")
	for _, v := range pairs {
		PrintVolumes(v)
	}
}

// Print in the terminal the volumes
func PrintVolumes(v map[string]interface{}) {
	fmt.Print("DEPLOYMENT -> ")
	fmt.Println(v["deploymentName"].(string))
	fmt.Print("VOLUME -> ")
	fmt.Println(v["volumeName"].(string))
	fmt.Print("DATA TYPE -> ")
	fmt.Println(v["dataTypeFrom"].(string))
	fmt.Print("SIZE -> ")
	fmt.Println(v["sizeFrom"].(string))
	fmt.Println("-------*------")
}

// TODO Show the migration or the process of the migration.
func ShowMigration(PathData string) {
	var pairs []map[string]interface{}

	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	fmt.Println("STATUS OF MIGRATION")
	fmt.Println("--------------------")
	for _, p := range pairs {
		fmt.Println("")
		fmt.Println("BACK UP")
		fmt.Println("--------------------")
		showRestic(p["deploymentName"].(string))
		fmt.Println("")
		fmt.Println("RECOVERY")
		fmt.Println("--------------------")
		showRecovery(p["deploymentName"].(string))
		fmt.Println("******************************************")
	}
}

//TODO
func showRestic(deployment string) {

	// DESCRIPTION ->
	fmt.Println("TODO " + deployment)
}

// TODO
func showRecovery(deployment string) {
	fmt.Println("TODO " + deployment)
}

