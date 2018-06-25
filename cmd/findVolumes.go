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
	"volume2volume/pkg/confObject"
	"volume2volume/pkg/utils"
)

// findVolumesCmd represents the findVolumes command
var findVolumesCmd = &cobra.Command{
	Use:   "findVolumes",
	Short: "Identify volumes",
	Long: `Identify the volumes in the two clusters and make the pairs`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("findVolumes called")
		FindAllVolumes(cmd, args)
		app.PairsVolumesByName(PathData, PathTemplate, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
		//fmt.Println(a[0]["podName"])
		//fmt.Println("---")
		//fmt.Println(b[0]["podName"])
	},
}

func init() {
	RootCmd.AddCommand(findVolumesCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findVolumesCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findVolumesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func FindAllVolumes(cmd *cobra.Command, args []string){
	//var PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo string
	//var ObjectsOc []string

	PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
	UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc =
		utils.GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
		ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)

	conf1 := confObject.ConfObject{PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
		UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc}
	fmt.Println("conf1 -->")
	fmt.Println(conf1)

	app.FindVolumes("ClusterFrom", PathTemplate, PathData, ClusterFrom, ClusterTo,
		ProjectTo, ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo ,ObjectsOc)

	app.FindVolumes("ClusterTo", PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
	 ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo ,ObjectsOc)
}
