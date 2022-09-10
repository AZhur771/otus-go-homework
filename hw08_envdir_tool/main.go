package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatalf("Failed to read env from directory: %v", err)
	}

	returnCode := RunCmd(args[2:], env)
	os.Exit(returnCode)
}
