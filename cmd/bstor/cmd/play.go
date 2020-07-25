package cmd

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/vistrcm/bstor/pgp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(playCmd)
}

var playCmd = &cobra.Command{
	Use:   "play",
	Short: "play with bstor",
	Long:  `dev only. Play with commandsm`,
	Run: func(cmd *cobra.Command, args []string) {
		play()
	},
}

func play() {

	g, err := pgp.New()
	if err != nil {
		panic(err)
	}

	publicKeyRing, err := g.GetKeyring("2496B1F0F2FD90F50FF574D548F22A464791F054")
	if err != nil {
		panic(err)
	}

	binMessage := crypto.NewPlainMessageFromString("plain text")

	pgpMessage, err := publicKeyRing.Encrypt(binMessage, nil)
	if err != nil {
		panic(err)
	}

	armored, err := pgpMessage.GetArmored()
	if err != nil {
		panic(err)
	}


	fmt.Println(armored)


}
