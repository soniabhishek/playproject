package commands

import (
	"fmt"
	"playment/source"
)

func ThrowError(err string) {
	source.Write(fmt.Sprintf("ERR: %s\n", err))
}
func ThrowSuccess(succ string) {
	source.Write(succ + "\n")
}

//function to be used prining pwd or in future along with username
func WriteOut() {
	if source.WriteSource == "os" {
		source.Write(fmt.Sprintf("$ %s:", Pwd))
	}
}
