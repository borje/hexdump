package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var infile = flag.String("input", "", "The file to read")

func readFromStdin() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Size() > 0 || fi.Mode()&os.ModeNamedPipe != 0
}

func main() {
	flag.Parse()
	var f *os.File
	if readFromStdin() {
		f = os.Stdin
	} else {
		var err error
		f, err = os.Open(*infile)
		if err != nil {
			log.Fatal("Failed to open ", *infile)
		}
	}
	dumper := hex.Dumper(os.Stdout)
	io.Copy(dumper, f)
	f.Close()
}
