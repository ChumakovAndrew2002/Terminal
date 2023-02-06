package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

type Terminal struct {
	CurrentPath      string
	CurrentFiles     []string
	CurrentDirs      []string
	terminalCommands map[string]Command
}

func NewTerminal() *Terminal {

	m := make(map[string]Command)

	terminal := &Terminal{
		CurrentPath:      "",
		CurrentFiles:     nil,
		CurrentDirs:      nil,
		terminalCommands: m,
	}

	//terminal.CurrentPath = "/"
	var err error
	terminal.CurrentPath, err = os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	terminal.terminalCommands["ls"] = &Ls{}
	terminal.terminalCommands["cd"] = &Cd{}
	terminal.terminalCommands["ln"] = &Ln{}
	terminal.terminalCommands["mkdir"] = &Mkdir{}
	terminal.terminalCommands["cp"] = &Cp{}
	terminal.terminalCommands["pwd"] = &Pwd{}
	terminal.terminalCommands["nano"] = &Nano{}
	terminal.terminalCommands["rm"] = &Rm{}
	terminal.terminalCommands["mv"] = &Mv{}
	terminal.terminalCommands["cat"] = &Cat{}
	terminal.terminalCommands["clear"] = &Clear{}
	terminal.terminalCommands["help"] = &Help{}

	return terminal
}

func (t *Terminal) GetFilesAndDirs() error {
	t.CurrentFiles = nil
	t.CurrentDirs = nil

	data, err := os.ReadDir(t.CurrentPath)
	if err != nil {
		return errors.Wrap(err, "directory are not exist to read dir, try again")
	}

	for _, datum := range data {
		if datum.IsDir() {
			t.CurrentDirs = append(t.CurrentDirs, datum.Name())
			continue
		}

		t.CurrentFiles = append(t.CurrentFiles, datum.Name())
	}

	return nil
}

func (t *Terminal) RetrieveTwoDirs() {
	dirs := ParsePath(t.CurrentPath)

	switch len(dirs) {
	case 0:
		fmt.Print("$[] ")
	case 1:
		fmt.Print("$[" + dirs[0] + "] ")
	default:
		fmt.Print("$[" + dirs[len(dirs)-2] + "/" + dirs[len(dirs)-1] + "] ")
	}
}
