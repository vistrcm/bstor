package cmd

import (
	"fmt"
	"github.com/vistrcm/bstor/pgp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt source ... target",
	Short: "encrypt files",
	Long:  `encrypt source to the target`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		s := args[0]
		t := args[1]

		encCopy(s, t)

	},
}

//encCopy encrypt files and copy
func encCopy(s string, t string) {
	// TODO: check src existence
	// TODO: check that target is ok
	fmt.Printf("Encrypt!: %+v -> %+v\n", s, t)

	g, err := pgp.New()
	if err != nil {
		panic(err)
	}

	publicKeyRing, err := g.GetKeyring("2496B1F0F2FD90F50FF574D548F22A464791F054")

	if err := pgp.EncryptFile(s, t, publicKeyRing); err != nil{
		panic(err)
	}
}