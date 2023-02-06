package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

type Ln struct {
	NameCommand string
	Arguments   []string
}

func (l *Ln) Execute(str string, t *Terminal) error {
	l.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	l.Arguments = arguments

	if flags != nil {
		fmt.Println("this does not an argument, rewrite command")
		return nil
	}

	if len(arguments) == 2 {
		err := os.Symlink(arguments[0], arguments[1])
		if err != nil {
			return errors.Wrap(err, "error when creating a link")
		}
	}

	fmt.Println("incorrect quantity of arguments")

	return nil
}

func (l *Ln) GetHelp(str string) {

}
