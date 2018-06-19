package app

import (
	"volume2volume/pkg/utils"
)

//It is used for initialize the clusters.
// Initialize the cluster for use stash (restic, recovery objects).
// Create the secrets for use the cloud storage (s3, minio, ....)
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
		utils.CreateObject("./templates/stash/stash-openshift.yaml")
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


