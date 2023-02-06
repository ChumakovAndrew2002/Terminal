package main

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"runtime"
)

type Clear struct {
}

func runCmd(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func (c *Clear) GetHelp(str string) {

}

func (c *Clear) Execute(str string, t *Terminal) error {
	switch runtime.GOOS {
	case "darwin":
		err := runCmd("clear")
		if err != nil {
			return errors.Wrap(err, "failed command run, try again")
		}
	case "linux":
		err := runCmd("clear")
		if err != nil {
			return errors.Wrap(err, "failed command run, try again")
		}
	case "windows":
		err := runCmd("cmd", "/c", "cls")
		if err != nil {
			return errors.Wrap(err, "failed command run, try again")
		}
	default:
		err := runCmd("clear")
		if err != nil {
			return errors.Wrap(err, "failed command run, try again")
		}
	}
	return nil
}
