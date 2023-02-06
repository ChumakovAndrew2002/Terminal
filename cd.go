package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

type Cd struct {
	NameCommand string
	Arguments   []string
}

func (c *Cd) Execute(str string, t *Terminal) error {
	c.NameCommand = GetNameCommand(str)
	arguments, flags := ParseRow(str)
	c.Arguments = arguments

	if len(arguments) > 1 {
		fmt.Println("too many arguments")
		return nil
	}

	if flags != nil {
		fmt.Println("this command dont have any flags, rewrite commnand")
		return nil
	}

	if arguments == nil {
		fmt.Println("not enough arguments")
		return nil
	}

	if arguments[0] == ".." {
		var newParts []string
		if t.CurrentPath == "/" {
			fmt.Println("you are in the root folder")
			return nil
		}

		parts := ParsePath(t.CurrentPath)
		newParts = append(newParts, parts[:len(parts)-1]...)
		parts = newParts
		newParts = nil
		t.CurrentPath = ""

		for _, com := range parts {
			t.CurrentPath = t.CurrentPath + "/" + com
		}

		if t.CurrentPath == "" {
			t.CurrentPath = "/"
		}

		return nil
	}

	dirs := strings.Split(arguments[0], "/")
	for _, dir := range dirs {
		err := t.GetFilesAndDirs()
		if err != nil {
			return errors.Wrap(err, "non-existent folder, try again")
		}

		for i, directory := range t.CurrentDirs {
			if directory == dir {
				if t.CurrentPath == "/" {
					t.CurrentPath = t.CurrentPath + dir
					break
				}

				t.CurrentPath = t.CurrentPath + "/" + dir
				break
			}

			if i == len(t.CurrentDirs)-1 {
				fmt.Println("a folder with this name does not exist: " + dir)
				return nil
			}
		}
	}

	return nil
}

func (c *Cd) GetHelp(str string) {

}

func ParsePath(path string) []string {
	var newParts []string
	partsOfPaths := strings.Split(path, "/")

	newParts = partsOfPaths

	if partsOfPaths[0] == "" {
		newParts = partsOfPaths[1:]
	}

	return newParts
}
