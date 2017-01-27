package main

import (
	"flag"

	"github.com/soniabhishek/playproject/source"
	"github.com/soniabhishek/playproject/command"
)

//this init function is used for parsing flag if we provide it with a file name
func init() {
	sourceFlag := flag.String("filename", "", "if input is through source file")
	flag.Parse()
	//this will initialize the reader and writer in source pkg
	source.Input(*sourceFlag)
	if *sourceFlag == "" {
		commands.WriteOut()
	}
}
