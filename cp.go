package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Cp struct {
	NameCommand string
	Arguments   []string
}

func (c *Cp) Execute(str string, t *Terminal) error {
	c.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	c.Arguments = arguments

	if flags != nil {
		fmt.Println("this does not an argument, rewrite command")
		return nil
	}

	if len(arguments) == 2 {
		path1 := arguments[0]
		path2 := arguments[1]
		if strings.HasPrefix(path1, "./") {
			name := strings.TrimPrefix(path1, "./")
			path1 = t.CurrentPath + "/" + name
		}
		if strings.HasPrefix(path2, "./") {
			name := strings.TrimPrefix(path2, "./")
			path2 = t.CurrentPath + "/" + name
		}

		data, err := os.ReadFile(path1)

		if err != nil {
			return errors.Wrap(err, "incorrect file path to read from file, try again")
		}

		err = os.WriteFile(path2, data, 0644)
		if err != nil {
			return errors.Wrap(err, "incorrect file path to write to file, try again")
		}

		return nil
	}

	fmt.Println("incorrect quantity of arguments")
	return nil
}

func (c *Cp) GetHelp(str string) {

}
