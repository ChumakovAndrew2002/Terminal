package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type Ls struct {
	NameCommand string
	Flags       []string
}

func (l *Ls) GetHelp(str string) {
	fmt.Println("help")
}

func (l *Ls) Execute(str string, t *Terminal) error {
	l.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	l.Flags = flags

	if arguments != nil {
		fmt.Println("this does not a flag, rewrite command")
		return nil
	}

	if len(flags) > 2 {
		fmt.Println("too many flags for that command")
		return nil
	}

	err := t.GetFilesAndDirs()

	if err != nil {
		return errors.Wrap(err, "non-existent folder, try again")
	}

	if len(l.Flags) != 0 {
		var ok bool

		for _, flag := range l.Flags {
			switch flag {
			case "f":
				ok = true
				for _, file := range t.CurrentFiles {
					fmt.Println(file)
				}

			case "d":
				ok = true
				for _, directory := range t.CurrentDirs {
					fmt.Println(directory)
				}
			}
		}

		if ok {
			return nil
		}

		fmt.Println("Undefined flag, try again")
		return nil
	}

	for _, directory := range t.CurrentDirs {
		fmt.Println(directory)
	}

	for _, file := range t.CurrentFiles {
		fmt.Println("    ", file)
	}

	return nil
}
