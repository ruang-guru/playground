package cmd

import (
	"github.com/ruang-guru/playground/cli/assignmentmanager"
	"github.com/spf13/cobra"
)

var assignmentFile string
var graderEmail string
var graderPassword string

var assignmentManagerCmd = &cobra.Command{
	Use:   "assignmentmanager",
	Short: "Manage assignments",
	Run: func(cmd *cobra.Command, args []string) {
		if err := assignmentmanager.ManageAssignmentsFromJsonFile(assignmentFile, graderEmail, graderPassword); err != nil {
			panic(err.Error())
		}
	},
}

func init() {
	assignmentManagerCmd.Flags().StringVarP(&assignmentFile, "file", "f", "", "path to json file containing list of assignments and their detail")
	assignmentManagerCmd.Flags().StringVarP(&graderEmail, "email", "", "", "email of grader teacher account")
	assignmentManagerCmd.Flags().StringVarP(&graderPassword, "password", "", "", "password of grader teacher account")

	assignmentManagerCmd.MarkFlagRequired("file")
	assignmentManagerCmd.MarkFlagRequired("email")
	assignmentManagerCmd.MarkFlagRequired("password")
	rootCmd.AddCommand(assignmentManagerCmd)
}
