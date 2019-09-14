// +build darwin

/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List downloaded kubectl binaries",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving installed versions of kubectl")

		err := ListKubectlBinaries()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ListKubectlBinaries - List available installed kubectl versions
func ListKubectlBinaries() error {

	// Get user home directory path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Read kubemngr directory
	kubectl, err := ioutil.ReadDir(homeDir + "/.kubemngr/")
	if err != nil {
		log.Fatal(err)
	}

	// iterate and print files inside the kubemngr directory
	for _, files := range kubectl {
		file := files.Name()
		version := strings.Replace(file, "kubectl-", "", -1)
		fmt.Println(version)
	}

	return nil
}
