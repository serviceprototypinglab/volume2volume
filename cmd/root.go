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
	"os"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ClusterFrom  string
	ClusterTo    string
	ProjectFrom  string
	ProjectTo    string
	PathTemplate string
	PathData     string
	UsernameFrom string
	UsernameTo   string
	PasswordFrom string
	PasswordTo   string
	cfgFile      string
)
var ObjectsOc []string
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "volume2volume",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.volume2volume.yaml)")


	RootCmd.PersistentFlags().StringVarP(&ClusterFrom, "clusterFrom", "", "", "Cluster where is the project that you want to migrate")
	RootCmd.PersistentFlags().StringVarP(&ClusterTo, "clusterTo", "", "", "Cluster where you want to migrate the project")
	RootCmd.PersistentFlags().StringVarP(&ProjectFrom, "projectFrom", "", "", "name of the old Openshift project")
	RootCmd.PersistentFlags().StringVarP(&ProjectTo, "projectTo", "", "", "name of the new Openshift project")
	RootCmd.PersistentFlags().StringVarP(&UsernameFrom, "usernameFrom", "", "", "username in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&UsernameTo, "usernameTo", "", "", "username in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PasswordFrom, "passwordFrom", "", "", "password in the cluster From")
	RootCmd.PersistentFlags().StringVarP(&PasswordTo, "passwordTo", "", "", "password in the cluster To")
	RootCmd.PersistentFlags().StringVarP(&PathTemplate, "pathTemplate","","", "path where export the templates")
	RootCmd.PersistentFlags().StringVarP(&PathData, "pathData","", "", "path where export the volumes")
	defaultValue := []string{""}
	RootCmd.PersistentFlags().StringArrayVarP(&ObjectsOc, "objects", "o", defaultValue, "list of objects to export" )

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".volume2volume" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".volume2volume")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}




