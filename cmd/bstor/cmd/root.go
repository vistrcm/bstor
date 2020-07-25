package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bstor",
	Short: "bstor command to store encrypted files",
	Long:  "a way to encrypt files and copy to external storage",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("hello")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
