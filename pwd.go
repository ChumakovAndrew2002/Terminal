package main

import "fmt"

type Pwd struct {
}

func (p *Pwd) Execute(str string, t *Terminal) error {
	fmt.Println(t.CurrentPath)
	return nil
}

func (p *Pwd) GetHelp(str string) {

}
