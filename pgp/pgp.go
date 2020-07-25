package pgp

import (
	"fmt"
	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"io/ioutil"
)

func EncryptFile(src, dst string, publicKeyRing *crypto.KeyRing) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("error opening source file %q: %w", src, err)
	}

	binMessage := crypto.NewPlainMessage(data)
	pgpMessage, err := publicKeyRing.Encrypt(binMessage, nil)
	if err != nil {
		return fmt.Errorf("encryption error: %w", err)
	}
	return ioutil.WriteFile(dst, pgpMessage.GetBinary(), 0600)
}
