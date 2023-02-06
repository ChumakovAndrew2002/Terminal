package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Cat struct {
	NameCommand string
	Arguments   []string
}

func (c *Cat) Execute(str string, t *Terminal) error {
	c.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	c.Arguments = arguments

	if flags != nil {
		fmt.Println("this does not an argument, rewrite command")
		return nil
	}

	if len(arguments) == 1 {
		if strings.Contains(arguments[0], "./") {
			name := strings.TrimPrefix(arguments[0], "./")
			arguments[0] = t.CurrentPath + "/" + name
		}

		data, err := os.ReadFile(arguments[0])
		if err != nil {
			return errors.Wrap(err, "incorrect path to cat, TRY AGAIN")
		}

		fmt.Println(string(data))

		return nil
	}
	fmt.Println("Wrong quantity of arguments")
	return nil
}

func (c *Cat) GetHelp(str string) {

}
