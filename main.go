package main

import (
	"github.com/soniabhishek/playproject/command"
	"github.com/soniabhishek/playproject/source"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	for source.Scanner.Scan() {

		//scan a new line and send it to comand variable to check command exist or not
		input := strings.Split(strings.Trim(source.Scanner.Text(), " "), " ")
		command := input[0]

		//blank line will result into new scanning again
		if command == "" {
			//writing utility like in console which describes pwd in future can use username
			commands.WriteOut()
			continue
		}

		//check if the command is valid otherwise throws error
		if !commands.Commands[command] {
			commands.ThrowError("CANNOT RECOGNIZE INPUT.")

		} else {
			//if a valid command than it will result in its execution
			commands.Execute(command, input[1:], nil)

		}
		//if using console than repeting prining of pwd
		commands.WriteOut()
	}
}
