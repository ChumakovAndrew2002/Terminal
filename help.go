package main

import "fmt"

type Help struct {
	Arguments    []string
	HelpCommands map[string]string
}

func (h *Help) Execute(str string, t *Terminal) error {
	arguments, flags := ParseRow(str)
	h.Arguments = arguments
	h.InitHelpCommands()

	if flags != nil {
		fmt.Println("this command doesnt use flags, rewrite command")
		return nil
	}

	if len(arguments) == 0 {
		fmt.Print(`
Usage:
      If you want  to start from current directory ,write "./"
      Write arguments and flags separated by spaces and write before the flags -

	command arguments -flags
       
      
Possible commands:
            ln       Create file's link
		help     Prints detailed information about application
		clear    Clears terminal screen
		cd       Changes current directory
		ls       Retrieves all files and directories in current path
		mv       Moves file or directory to the new location
		cat      Prints file content
		cp       Creates file copy
		rm       Removes file or directory
		nano     Creates and writes some data to the file
		mkdir    Creates new directory
		pwd      Prints current path

You can also see detailed information about each command:
		help command
`)
		return nil
	}

	if len(arguments) == 1 {
		fmt.Println(arguments[0])
		val, ok := h.HelpCommands[arguments[0]]
		if !ok {
			fmt.Println("command not found, try again")
			return nil
		}

		fmt.Println(val)

		return nil
	}

	fmt.Println("wrong quantity of arguments, rewrite command")
	return nil
}

func (h *Help) GetHelp(str string) {
}

func (h *Help) InitHelpCommands() {
	h.HelpCommands = make(map[string]string)

	h.HelpCommands["ln"] = `
Create file's link
How to use: write ln and then write which file or directory you want to a have a link and then write command where you want to put the link
Example : $[] ln path/1 path/2
`
	h.HelpCommands["clear"] = `
Clears terminal screen
How to use: write clear without any arguments
Example : $[] clear
`
	h.HelpCommands["cd"] = `
Changes current directory
How to use: write cd and then write file or directory with which one do you want to work
Example : $[] cd directory
`
	h.HelpCommands["ls"] = `
Retrieves all files and directories in current path
How to use: write ls, and if you want write any flags: 
-f - write only files in current directory
-d - write only directories in current directory
If you want use more than one flag, write them through space "-f -d"
Example : $[] ls -f
`
	h.HelpCommands["pwd"] = `
Prints current path
How to use: write pwd without any arguments
Example : $[] pwd
`
	h.HelpCommands["cp"] = `
Creates file copy
How to use: write cp and then write two paths , the first path is to what you want to copy 
and the second path is where you want to copy

Example : $[] cp /first/path /second/path
`
	h.HelpCommands["mv"] = `
Moves file or directory to the new location
How to use: write mv and then write two paths , the first path is that you want to move
and second path is where you to put this

Example : $[] mv first/path second/path
`
	h.HelpCommands["cat"] = `
Prints file content
How to use: write cat and than write path to thing that you want to print in terminal
Example : $[] cat first/path
`
	h.HelpCommands["rm"] = `
Removes file or directory
How to use: write rm and than write path to thing that you want remove 
Example : $[] rm first/path
`
	h.HelpCommands["nano"] = `
Creates and writes some data to the file
How to use: write the name of the file in current directory
Example : $[] nano name
`
	h.HelpCommands["mkdir"] = `
Creates new directory
How to use: write mkdir , then write name of new directory and then write path where you want to make new directory
Example : $[] mkdir name path/where/newDirectory
`
}
