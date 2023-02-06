package main

import "strings"

func ScreeningSymbols(str string) []string {
	var symbols []string
	var newCommand string
	var k int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == `\` && k == i {
			newCommand = newCommand + string(str[i])
			continue
		}

		if string(str[i]) == `\` && i != len(str)-1 {
			if string(str[i+1]) == `-` {
				newCommand = newCommand + string(str[i])
			}
			k = i + 1
			continue
		}

		if string(str[i]) == " " && k != i {
			symbols = append(symbols, newCommand)
			newCommand = ""
			continue
		}

		if i == len(str)-1 {
			newCommand = newCommand + string(str[i])
			symbols = append(symbols, newCommand)
			newCommand = ""
			continue
		}

		newCommand = newCommand + string(str[i])
	}

	return symbols
}

func ParseRow(str string) ([]string, []string) {
	parts := ScreeningSymbols(str)
	var arguments, flags []string

	for _, part := range parts {
		if part == "" {
			continue
		}

		if string(part[0]) == "-" {
			if len(part) > 1 {
				part = strings.TrimPrefix(part, "-")
				flags = append(flags, part)
			}
		} else {
			if string(part[0]) == `\` {
				part = strings.TrimPrefix(part, `\`)
			}
			arguments = append(arguments, part)
		}

	}

	return arguments, flags
}
