package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	// Place your code here.
	files, err := ioutil.ReadDir("testdata")

	var filenames []string

	for _, file := range files {
		filename := file.Name()
		if filename != "input.txt" {
			filenames = append(filenames, file.Name())
		}
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, filename := range filenames {
		filename := filename
		t.Run(fmt.Sprintf("Test - %s", filename), func(t *testing.T) {
			tmpFile, err := os.CreateTemp("", "tmp.*.txt")
			if err != nil {
				log.Fatalf("Some error occurred: %v", err)
			}

			re := regexp.MustCompile(`\d+`)
			nums := re.FindAllString(filename, -1)
			offset, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatalf("Some error occurred: %v", err)
			}

			limit, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatalf("Some error occurred: %v", err)
			}

			err = Copy(filepath.Join("testdata", "input.txt"), tmpFile.Name(), int64(offset), int64(limit))
			require.NoError(t, err)

			expected, err := ioutil.ReadFile(filepath.Join("testdata", filename))
			if err != nil {
				log.Fatalf("Some error occurred: %v", err)
			}

			result, err := ioutil.ReadFile(tmpFile.Name())
			if err != nil {
				log.Fatalf("Some error occurred: %v", err)
			}

			require.Equal(t, string(expected), string(result))
		})
	}
}
