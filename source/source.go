package source

import (
	"bufio"
	"log"
	"os"
)

//this will be used as a flag to determine whether the source is file or console or any other case in future
var WriteSource string

var Scanner *bufio.Scanner
var Writer *bufio.Writer

func Input(source string) {
	//called from the init function this sets the reader and writer

	if source != "" {
		readFile, err := os.OpenFile(source, os.O_RDONLY, os.ModeExclusive)
		if err != nil {
			log.Fatal(err)
		}

		Scanner = bufio.NewScanner(readFile)
		//the output file would be like this
		writeFile, err := os.Create(source + "_output.txt")
		if err != nil {
			log.Fatal(err)
		}

		Writer = bufio.NewWriter(writeFile)
		WriteSource = "file"

	} else {

		Scanner = bufio.NewScanner(os.Stdin)
		Writer = bufio.NewWriter(os.Stdout)
		WriteSource = "os"
	}
}

//this will write in the initialized writer
func Write(content string) {
	Writer.WriteString(content)
	Writer.Flush()
}
