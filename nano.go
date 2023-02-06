package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Nano struct {
}

func (n *Nano) Execute(str string, t *Terminal) error {

	path := t.CurrentPath + "/" + str

	err := t.GetFilesAndDirs()
	if err != nil {
		return errors.Wrap(err, "failed command GetFilesAndDirs, try again")
	}

	for _, file := range t.CurrentFiles {
		if file == str {
			newC := &Clear{}
			err = newC.Execute(str, t)
			if err != nil {
				errors.Wrap(err, "failed command clear, try again")
			}
			fmt.Println("Write your data in file, if you want to stop writting in file write \"....\"")

			err := t.WriteDataInFile(path)
			return errors.Wrap(err, "failed writing data in file, try again")
		}
	}

	err = t.MakeFile(path)
	if err != nil {
		return errors.Wrap(err, "failed making file, try again")
	}

	newC := &Clear{}
	err = newC.Execute(str, t)
	if err != nil {
		errors.Wrap(err, "failed command clear, try again")
	}
	fmt.Println("Write your data in file, if you want to stop writting in file write \"....\"")

	err = t.WriteDataInFile(path)
	if err != nil {
		return errors.Wrap(err, "failed writing data in file, try again")
	}

	return nil
}

func (t *Terminal) MakeFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create file, incorrect path,try again")
	}
	f.Close()

	return nil
}

func (t *Terminal) WriteDataInFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return errors.Wrap(err, "incorrect path, didn`t reading file, try again")
	}

	var stopWord bool
	var newData []string

	newData = append(newData, string(data))
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		return errors.Wrap(err, "failed opening file, wrong path, try again")
		os.Exit(1)
	}

	for !stopWord {
		str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		str = strings.Trim(str, "\n")

		if strings.HasSuffix(str, "....") {
			stopWord = strings.HasSuffix(str, "....")
			break
		}

		w := bufio.NewWriter(file) // Создаем новый объект Writer

		_, err = file.WriteString("\n" + str)
		if err != nil {
			return errors.Wrap(err, "failed writing string in file, try again")
		}

		err = w.Flush()
		if err != nil {
			return errors.Wrap(err, "flush")
		}

		//stopWord = strings.HasSuffix(str, "....")
	}

	defer file.Close()
	return nil
}

func (n *Nano) GetHelp(str string) {

}
