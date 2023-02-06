package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Rm struct {
	NameCommand string
	Arguments   []string
}

func (r *Rm) Execute(str string, t *Terminal) error {
	r.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	r.Arguments = arguments

	if flags != nil {
		fmt.Println("this does not an argument, rewrite command")
		return nil
	}

	if len(arguments) > 1 {
		fmt.Println("too many arguments for that command")
		return nil
	}

	if arguments == nil {
		fmt.Println("not enough arguments")
		return nil
	}

	path := arguments[0]

	err := t.GetFilesAndDirs()
	if err != nil {
		return errors.Wrap(err, "failed getting files and directories, try again")
	}

	if strings.HasPrefix(arguments[0], "./") {
		name := strings.TrimPrefix(arguments[0], "./")
		path = t.CurrentPath + "/" + name
	}

	err = os.Remove(path)
	if err != nil {
		return errors.Wrap(err, "failed remove file, incorrect path, try again")
	}

	return nil
}

func (r *Rm) GetHelp(str string) {

}
