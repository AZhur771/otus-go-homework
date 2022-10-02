package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	release   = "UNKNOWN"
	buildDate = "UNKNOWN"
	gitHash   = "UNKNOWN"
)

func getGitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--verify", "HEAD")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(out), "\n"), nil
}

func printVersion() {
	if gitHash == "UNKNOWN" {
		if hash, err := getGitHash(); err != nil {
			fmt.Printf("error while get hash of the last commit: %v\n", err)
		} else {
			gitHash = hash
		}
	}

	if err := json.NewEncoder(os.Stdout).Encode(struct {
		Release   string
		BuildDate string
		GitHash   string
	}{
		Release:   release,
		BuildDate: buildDate,
		GitHash:   gitHash,
	}); err != nil {
		fmt.Printf("error while decode version info: %v\n", err)
	}
}
