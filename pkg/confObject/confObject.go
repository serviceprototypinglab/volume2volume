	// ConfObject package.
	// Configuration objects for the migration of the data.
	// This configuration can be given in a config file or using flags.
	// It is not used now.
package confObject

/*
	- PathTemplate = Path where volume2volume can find all the templates (recovery, restic, stats, stash, ...)

	- PathData = Path where you can find the metadata of the volumes.

	- ClusterFrom = URL of the cluster From.

	- ClusterTo = URL of the cluster To.

	- ProjectFrom = Name of the project in cluster From.

	- ProjectTo= Name of the project in cluster To.

	- UsernameFrom = Name for the user used in cluster From.

	- UsernameTo = Name for the user used in cluster To.

	- PasswordFrom = Password for the user used in cluster From.

	- PasswordTo = Password for the user used in cluster To.

	- ObjectsOc = Types of the objects that you want to migrate (deployments, ...).
*/
type ConfObject struct {
	PathTemplate string
	PathData     string
	ClusterFrom  string
	ClusterTo    string
	ProjectTo    string
	ProjectFrom  string
	UsernameTo   string
	UsernameFrom string
	PasswordFrom string
	PasswordTo   string
	ObjectsOc    []string
}

type confCluster struct {
	Cluster  string
	Project  string
	Username string
	Password string
}