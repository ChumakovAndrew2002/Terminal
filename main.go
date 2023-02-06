package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strings"
)

type Command interface {
	Execute(str string, t *Terminal) error
	GetHelp(str string)
}

func main() {
	term := NewTerminal()

	for {
		term.RetrieveTwoDirs()
		str, err := Reader()
		if err != nil {
			log.Fatal(err)
		}

		name := GetNameCommand(str)
		str = strings.TrimPrefix(str, name+" ")
		str = strings.TrimPrefix(str, name)

		val, ok := term.terminalCommands[name]
		if !ok {
			fmt.Println("command not found")
			continue
		}

		err = val.Execute(str, term)
		if err != nil {
			fmt.Println(err)
			err = val.Execute(str, term)
		}
	}
}

func Reader() (string, error) {
	str, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", errors.Wrap(err, "failed new reader, try again")
	}
	str = strings.Trim(str, "\n")
	return str, nil
}

func GetNameCommand(command string) string {
	var nameCommand string
	values := strings.Split(command, " ")
	nameCommand = values[0]
	return nameCommand
}
