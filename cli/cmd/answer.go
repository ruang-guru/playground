package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Answer utilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("answer called")
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// answerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// answerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
