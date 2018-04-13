package app

import (
	"fmt"
	"volume2volume/pkg/utils"
)

//TODO Show all the volumes pairs and statistics in a nice way.
func ShowVolumes(PathData string) {
	var pairs []map[string]interface{}

	pairs = utils.ReadJsonArray(PathData + "/pairs/", "pairs")
	fmt.Println("PAIRS OF VOLUMES")
	fmt.Println("-------*------")
	for _, v := range pairs {
		PrintVolumes(v)
	}
}

func PrintVolumes(v map[string]interface{}) {
	fmt.Print("DEPLOYMENT -> ")
	fmt.Println(v["deploymentName"].(string))
	fmt.Print("VOLUME -> ")
	fmt.Println(v["volumeName"].(string))
	fmt.Println("-------*------")
}
