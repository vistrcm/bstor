package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of bstor",
	Long:  `All software has versions. This is bstors's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bstor -- HEAD")
	},
}
