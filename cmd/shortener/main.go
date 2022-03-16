package main

import (
	"log"
	"os"
)

func main() {
	if true {
		log.Println("Exit is here")
		os.Exit(0)
	} else {
		log.Println("Never")
	}
}
