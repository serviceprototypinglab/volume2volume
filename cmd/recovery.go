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
	"volume2volume/pkg/utils"
	"volume2volume/pkg/app"
)

// recoveryCmd represents the recovery command
var recoveryCmd = &cobra.Command{
	Use:   "recovery",
	Short: "Recovery all the volumes paired",
	Long: `Create the recovery objects to upload the data of the volumes. (in the new cluster)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recovery called")
		PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
			UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc =
			utils.GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
				ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
		app.Recovery(PathData)
	},
}

func init() {
	RootCmd.AddCommand(recoveryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recoveryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recoveryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
