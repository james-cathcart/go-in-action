package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

func init() {
	// change the logging device to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}