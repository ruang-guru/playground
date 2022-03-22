/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/ruang-guru/playground/cli/answerremover"
	"github.com/spf13/cobra"
)

var rootFolder string
var excludeFolders []string
var fileNames []string
var extensions []string

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove answer from all codes",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Removing all answer")
		if err := answerremover.RemoveAllAnswerBlocks(rootFolder,
			fileNames,
			extensions,
			excludeFolders); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	answerCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVarP(&rootFolder, "rootFolder", "r", "", "Root folder of the project")
	removeCmd.Flags().StringArrayVarP(&excludeFolders, "excludeFolders", "e", []string{"libs/tools/answerremover", "node_modules"}, "Exclude folders")
	removeCmd.Flags().StringArrayVar(&fileNames, "filenames", []string{"Dockerfile", "Makefile", "docker-compoose.yaml"}, "Extensions")
	removeCmd.Flags().StringArrayVar(&extensions, "extensions", []string{".ts", ".go", ".yaml", ".yml", ".js", ".css", ".tsx", ".proto", ".conf", ".md", ".html"}, "Extensions")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
