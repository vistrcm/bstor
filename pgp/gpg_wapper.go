// whole module is inspired by https://github.com/keybase/client/blob/master/go/libkb/gpg_cli.go
package pgp

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
)

//GpgCLI representation of GpgCLI.
// inspired by https://github.com/keybase/client/blob/master/go/libkb/gpg_cli.go
type GpgCLI struct {
	path string
}

//New() construct and return new GpgCLI.
func New() (GpgCLI, error) {
	path, err := findGpgPath()
	if err != nil {
		return GpgCLI{}, fmt.Errorf("error initializing new GpgCLI: %w", err)
	}

	return GpgCLI{path: path}, nil
}

//findGpgPath looks for gpg2, if not found for gpg and return an path for this executable.
func findGpgPath() (string, error) {
	prog, err := exec.LookPath("gpg2")

	if err != nil {
		prog, err = exec.LookPath("gpg")
	}

	if err != nil {
		return "", err
	}

	return prog, nil
}

type RunGpg2Res struct {
	Stdout []byte
	Stderr []byte
	Err    error
}

type readChanRes struct {
	data []byte
	err  error
}

func (rc readChanRes) result() ([]byte, error) {
	return rc.data, rc.err
}

//run command
func (g GpgCLI) run(args ...string) (res RunGpg2Res) {
	const cmdTimeout = 100 * time.Millisecond

	ctx, cancel := context.WithTimeout(context.Background(), cmdTimeout)
	defer cancel()

	// Always use --no-auto-check-trustdb to prevent gpg from refreshing trustdb.
	// details: https://github.com/keybase/client/blob/master/go/libkb/gpg_cli.go#L396
	args = append([]string{"--no-auto-check-trustdb"}, args...)

	cmd := exec.CommandContext(ctx, g.path, args...)

	var stdout, stderr io.ReadCloser

	if stdout, res.Err = cmd.StdoutPipe(); res.Err != nil {
		return
	}

	if stderr, res.Err = cmd.StderrPipe(); res.Err != nil {
		return
	}

	if res.Err = cmd.Start(); res.Err != nil {
		return
	}

	stdoutCh := make(chan readChanRes)
	stderrCh := make(chan readChanRes)

	go func() {
		data, err := ioutil.ReadAll(stdout)
		stdoutCh <- readChanRes{
			data: data,
			err:  err,
		}
	}()

	go func() {
		data, err := ioutil.ReadAll(stderr)
		stderrCh <- readChanRes{
			data: data,
			err:  err,
		}
	}()

	if res.Stderr, res.Err = (<-stderrCh).result(); res.Err != nil {
		return
	}

	if res.Stdout, res.Err = (<-stdoutCh).result(); res.Err != nil {
		return
	}

	if res.Err = cmd.Wait(); res.Err != nil {
		return
	}

	return res
}

//GetKey return pgp key for the `id`.
func (g GpgCLI) GetKey(id string) ([]byte, error) {
	res := g.run("--export", id)

	if res.Err != nil {
		return nil, fmt.Errorf("error getting key: %w", res.Err)
	}

	return res.Stdout, nil
}

//GetKeyring returns `crypto.KeyRing` of a specific key `id`
func (g GpgCLI) GetKeyring(id string) (*crypto.KeyRing, error) {
	pubkey, err := g.GetKey(id)
	if err != nil {
		return nil, fmt.Errorf("error getting key bytes for id %q: %w", id, err)
	}

	publicKeyObj, err := crypto.NewKey(pubkey)
	if err != nil {
		return nil, fmt.Errorf("error getting key for id %q: %w", id, err)
	}

	return crypto.NewKeyRing(publicKeyObj)
}
