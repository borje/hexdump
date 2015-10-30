package main

import (
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
)

var infile = flag.String("input", "", "The file to read")

func main() {
	flag.Parse()
	log.Println("file: ", *infile)

	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal("Failed to open ", *infile)
	}
	dumper := hex.Dumper(os.Stdout)
	io.Copy(dumper, f)
}
