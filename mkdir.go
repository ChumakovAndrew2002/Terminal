package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Mkdir struct {
	NameCommand string
	Arguments   []string
}

//todo сделать возможность создания папкки в папке для абсолютного пути

func (m *Mkdir) Execute(str string, t *Terminal) error {
	m.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	m.Arguments = arguments

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

	if strings.HasPrefix(arguments[0], "./") {
		name := strings.TrimPrefix(arguments[0], "./")

		err := MakeALLDirs(name, t)
		if err != nil {
			return errors.Wrap(err, "cannot make all dirs")
		}
	}

	err := os.Mkdir(arguments[0], 0777)
	if err != nil {
		return errors.Wrap(err, "incorrect path to make directory, try again")
	}

	return nil
}

func MakeALLDirs(str string, t *Terminal) error {
	names := strings.Split(str, "/")

	for _, name := range names {
		err := t.GetFilesAndDirs()
		if err != nil {
			return errors.Wrap(err, "non-existent folder, try again")
		}

		newPath := t.CurrentPath + "/" + name

		err = os.Mkdir(newPath, 0777)
		if err != nil {
			return errors.Wrap(err, "incorrect path to make directory, try again")
		}
		t.CurrentPath = newPath
	}
	return nil
}

func (m *Mkdir) GetHelp(str string) {

}

//func (t *Terminal) Mkdir(name, path string) error {
//	if path == "./" {
//		path = t.CurrentPath
//	}
//
//	newPath := path + "/" + name
//	err := os.Mkdir(newPath, 0777)
//	if err != nil {
//		return errors.Wrap(err, "incorrect path to make directory, try again")
//	}
//	return nil
//}
