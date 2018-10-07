package main

import (
	"bufio"
	"fmt"
	"github.com/applejan/dwgbreaker/brfile"
	"io/ioutil"
	"log"
	"os"
)

const (
	DATABASE = "C:\\ProgramData\\DwgBreaker\\info.br"
	REPLACER = [4]byte{00, 00, 00, 00}
)

type dwg struct {
	path     string
	filename string
	replc    string
}

func replace() {
	//TODO
}
func main() {
	file, err := ioutil.ReadFile("JG.dwg")
	if err != nil {
		log.Print(err)
	}
	re := bufio.NewReader(os.Stdin)
	brfile
	replc := file[:5]
	fmt.Print(string(replc))
}
