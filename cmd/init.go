// Copyright © 2018 Manuel Ramirez Lopez <ramz@zhaw.ch>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"volume2volume/pkg/app"
	"volume2volume/pkg/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the clusters",
	Long: `Initialize the cluster to get ready for stash objects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		//clusterFrom
		configureCluster("clusterFrom")
		//clusterTo
		configureCluster("clusterTo")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}



// Identify the version of the cluster and create the proper stash version.
func configureCluster(cluster string) {
	// TODO add support to the restic, recovery objects to the cluster and secrets

	fmt.Println(cluster)
	PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
		UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc =
		utils.GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
			ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)



	app.InitClusters(ClusterFrom, ClusterTo, ProjectFrom, ProjectTo,
		UsernameFrom, PasswordFrom, UsernameTo, PasswordTo)

}