package confObject

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