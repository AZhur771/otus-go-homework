package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Correctly run executor with command `date`", func(t *testing.T) {
		errorCode := RunCmd([]string{"date"}, Environment{})
		require.Equal(t, 0, errorCode)
	})

	t.Run("Fail to run executor with command `git something`", func(t *testing.T) {
		errorCode := RunCmd([]string{"git", "something"}, Environment{})
		require.Equal(t, 1, errorCode)
	})
}
