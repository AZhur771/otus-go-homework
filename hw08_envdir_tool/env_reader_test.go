package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Correctly read env variables from env dir", func(t *testing.T) {
		result := Environment{
			"BAR": EnvValue{
				Value:      "bar",
				NeedRemove: false,
			},
			"EMPTY": EnvValue{
				Value:      "",
				NeedRemove: true,
			},
			"FOO": EnvValue{
				Value:      "   foo\nwith new line",
				NeedRemove: false,
			},
			"HELLO": EnvValue{
				Value:      "\"hello\"",
				NeedRemove: false,
			},
			"UNSET": EnvValue{
				Value:      "",
				NeedRemove: true,
			},
		}

		env, err := ReadDir("./testdata/env")
		if err != nil {
			log.Fatal("Test went wrong")
		}

		for k, v := range env {
			resultV, ok := result[k]
			if !ok {
				log.Fatal("Test went wrong")
			}

			require.Equal(t, resultV.NeedRemove, v.NeedRemove)
			require.Equal(t, resultV.Value, v.Value)
		}
	})

	t.Run("Fail read env variables from env2 dir", func(t *testing.T) {
		dirName, err := ioutil.TempDir("", "test")
		defer os.RemoveAll(dirName)
		if err != nil {
			log.Fatal("Test went wrong")
		}

		tmpFile, err := ioutil.TempFile(dirName, "")
		if err != nil {
			log.Fatal("Test went wrong")
		}

		tmpFile.WriteString("foo=bar")

		_, err = ReadDir(dirName)
		require.ErrorIs(t, ErrInvalidFile, err)
	})
}
