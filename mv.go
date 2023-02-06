package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Mv struct {
	NameCommand string
	Arguments   []string
}

func (m *Mv) Execute(str string, t *Terminal) error {
	m.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	m.Arguments = arguments

	if flags != nil {
		fmt.Println("this doesnt have a flags, rewrite command")
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

		err = os.RemoveAll(path1)
		if err != nil {
			return errors.Wrap(err, "crashed deleting file, try again")
		}

		partsOfPath := strings.Split(path1, "/")
		name := partsOfPath[len(partsOfPath)-1]
		path2 = path2 + "/" + name

		f, err := os.Create(path2)
		if err != nil {
			return errors.Wrap(err, "failed creating file, try again")
		}

		err = f.Close()
		if err != nil {
			return errors.Wrap(err, "failed closing file, try again")
		}

		err = os.WriteFile(path2, data, 0644)
		if err != nil {
			return errors.Wrap(err, "incorrect file path to write to file, try again")
		}

		return nil
	}

	fmt.Println("wrong quantity of arguments")
	return nil
}

func (m *Mv) GetHelp(str string) {

}
