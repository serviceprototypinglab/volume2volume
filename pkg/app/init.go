package app

import (
	"volume2volume/pkg/utils"
)

func InitCluster(cluster, username, password string) {

	//stash := utils.ReadJson("templates/stash", "stash-openshift")
	//fmt.Println(stash)

	// Create stash in the cluster
	utils.LoginAdmin(cluster)
	utils.CreateObject("./templates/stash/stash-openshift.json")

	// Create secrets for backend.
	// s3
	utils.LoginCluster(cluster, username, password)
	utils.CreateSecret("S3")

	// Minio
	utils.CreateSecret("MINIO")

}
