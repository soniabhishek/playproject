package commands

import (
	"fmt"
	"strings"
	"github.com/soniabhishek/playproject/source"
)

var (
	Commands = map[string]bool{
		"pwd":     true,
		"cd":      true,
		"mkdir":   true,
		"rm":      true,
		"session": true,
	}

	mapper = map[string]execute{
		"pwd":   getPwd,
		"cd":    changeDirectory,
		"mkdir": makeDirectory,
		"rm":    removeDirectory,
	}
	Directories = []string{"/"}
	Pwd         = "/"
)

type execute func([]string, map[string]string) bool

func Execute(commandName string, arguments []string, flags map[string]string) {
	mapper[commandName](arguments, flags)
}

func isAbsolute(input byte) bool {
	if input == '/' {
		return true
	} else {
		return false
	}
}

func getPwd(arguments []string, flags map[string]string) bool {
	source.Write(Pwd + "\n")
	return true
}

func changeDirectory(arguments []string, flags map[string]string) bool {
	argument := arguments[0]
	if len(arguments) < 1 {
		ThrowError("Invalid Parameters")
	}
	if !isAbsolute(argument[0]) {
		argument = Pwd + argument
	}
	if argument[len(argument)-1] != '/' {
		argument = argument + "/"
	}

	if AlreadyExists(argument) {
		Pwd = argument
		ThrowSuccess("SUCC: REACHED")

	} else {
		ThrowError("Invalid Directory")
	}
	return true
}

func makeDirectory(arguments []string, flags map[string]string) bool {
	argument := arguments[0]
	if !MakeDirectory(argument) {
		ThrowError("Invalid Path")
	} else {
		ThrowSuccess("SUCC: CREATED")
	}

	return true
}

func removeDirectory(arguments []string, flags map[string]string) bool {
	argument := arguments[0]
	if len(arguments) < 1 {
		ThrowError("Invalid Parameters")
	}

	if !isAbsolute(argument[0]) {
		argument = Pwd + argument
	}
	if argument[len(argument)-1] != '/' {
		argument = argument + "/"
	}
	fmt.Println(Directories)
	if AlreadyExists(argument) {
		for x, v := range Directories {
			if v == argument {
				Directories = append(Directories[:x], Directories[x:]...)
				break
			}
		}
		fmt.Println(Directories)

		ThrowSuccess("SUCC: DELETED")
	} else {
		ThrowError("Invalid Object")
	}
	return true
}

func AlreadyExists(name string) bool {
	for _, x := range Directories {
		if x == name {
			return true
		}
	}
	return false
}
func CheckParent(name string) bool {

	all := strings.Split(name, "/")

	parent := ""
	for _, v := range all[:len(all)-1] {
		if v != " " {
			parent = parent + v + "/"
		}
	}

	if AlreadyExists(parent) {
		Directories = append(Directories, name+"/")
		return true
	}

	return false
}

func MakeDirectory(input string) bool {
	if len(input) < 1 {
		return false
	}
	if isAbsolute(input[0]) {
		if AlreadyExists(input + "/") {
			return false
		} else {
			if CheckParent(input) {
				return true
			}
		}
	} else {
		if AlreadyExists(Pwd + input + "/") {
			return false
		} else {
			if CheckParent(Pwd + input) {
				return true
			}
		}
	}
	return false
}
