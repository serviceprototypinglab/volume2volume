package app

import (
	"volume2volume/pkg/utils"
)

func InitClusters(clusterFrom, clusterTo, projectFrom, projectTo, usernameFrom,
	passwordFrom, usernameTo, passwordTo string) {

	//stash := utils.ReadJson("templates/stash", "stash-openshift")
	//fmt.Println(stash)

	// Create stash in the cluster
	utils.LoginAdmin(clusterFrom)
	utils.CreateObject("./templates/stash/stash-openshift.json")

	if clusterFrom != clusterTo {
		// Create stash in the cluster
		utils.LoginAdmin(clusterTo)
		utils.CreateObject("./templates/stash/stash-openshift.json")
	}

	// Create secrets for restic backend.
	// ClusterFrom
	// s3
	utils.LoginCluster(clusterFrom, usernameFrom, passwordFrom)
	utils.ChangeProject(projectFrom)
	utils.CreateSecret("s3")
	// minio
	utils.CreateSecret("minio")

	// Create secrets for restic backend.
	// ClusterTo
	// s3
	utils.LoginCluster(clusterTo, usernameTo, passwordTo)
	utils.ChangeProject(projectTo)
	utils.CreateSecret("s3")
	// minio
	utils.CreateSecret("minio")

}


