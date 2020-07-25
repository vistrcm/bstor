package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
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
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func addCommands() {
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(encryptCmd())
	rootCmd.AddCommand(playCmd())
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
