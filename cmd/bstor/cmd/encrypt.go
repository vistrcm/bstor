package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/vistrcm/bstor/pgp"
)

func encryptCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "encrypt source ... target",
		Short: "encrypt files",
		Long:  `encrypt source to the target`,
		Args:  cobra.ExactArgs(2), //nolint: gomnd
		Run: func(cmd *cobra.Command, args []string) {
			s := args[0]
			t := args[1]

			encCopy(s, t)
		},
	}
}

//encCopy encrypt files and copy.
func encCopy(s string, t string) {
	if err := checkSrc(s); err != nil {
		panic(err)
	}

	if err := checkDst(t); err != nil {
		panic(err)
	}

	fmt.Printf("Encrypt!: %+v -> %+v\n", s, t)

	g, err := pgp.New()
	if err != nil {
		panic(err)
	}

	publicKeyRing, err := g.GetKeyring("2496B1F0F2FD90F50FF574D548F22A464791F054")
	if err != nil {
		panic(err)
	}

	if err := pgp.EncryptFile(s, t, publicKeyRing); err != nil {
		panic(err)
	}
}

func checkDst(t string) error {
	//check if path exist
	_, err := os.Stat(path.Dir(t))
	return err
}

func checkSrc(s string) error {
	// check if file exist
	_, err := os.Stat(s)
	// error accessing file. One of the possibilities: os.ErrNotExist
	return err
}
