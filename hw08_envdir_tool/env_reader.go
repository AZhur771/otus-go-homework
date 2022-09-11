package main

import (
	"bytes"
	"errors"
	"os"
	"path"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

var ErrInvalidFile = errors.New("invalid file")

func getFileContent(b []byte) (string, error) {
	newLineIndex := bytes.IndexByte(b, 0x0A)
	if newLineIndex != -1 {
		b = b[0:newLineIndex]
	}
	b = bytes.ReplaceAll(b, []byte{0x00}, []byte{0x0A})
	str := string(b)
	str = strings.TrimRight(str, "\n\t ")

	if strings.Contains(str, "=") {
		return "", ErrInvalidFile
	}

	return str, nil
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	fileNames, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range fileNames {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		b, err := os.ReadFile(path.Join(dir, fileName))
		if err != nil {
			return nil, err
		}

		str, err := getFileContent(b)
		if err != nil {
			return nil, err
		}

		env[fileName] = EnvValue{
			Value:      str,
			NeedRemove: len(str) == 0,
		}
	}

	return env, nil
}
