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

// recoveryVolumeCmd represents the recoveryVolume command
var recoveryVolumeCmd = &cobra.Command{
	Use:   "recoveryVolume",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("recoveryVolume called")
		// TODO. Take names from configuration
		deploymentName := "deployment1"
		volumeName := "deployment1-storage"
		PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo, ProjectFrom,
			UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc =
			utils.GetAllValueReturn(PathTemplate, PathData, ClusterFrom, ClusterTo, ProjectTo,
				ProjectFrom, UsernameTo, UsernameFrom, PasswordFrom, PasswordTo, ObjectsOc)
		app.RecoveryVolume(PathData, deploymentName, volumeName)
	},
}

func init() {
	RootCmd.AddCommand(recoveryVolumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// recoveryVolumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// recoveryVolumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
