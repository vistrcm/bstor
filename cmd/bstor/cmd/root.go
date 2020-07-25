package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bstor",
		Short: "bstor command to store encrypted files",
		Long:  "a way to encrypt files and copy to external storage",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello")
		},
	}
}

//Execute main command.
func Execute() {
	rootCmd := rootCmd()
	addCommands(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		er(err)
	}
}

func addCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(encryptCmd())
	rootCmd.AddCommand(playCmd())
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
