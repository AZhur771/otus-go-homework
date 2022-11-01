package storage

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPqDuration(t *testing.T) {
	t.Run("Test scanning", func(t *testing.T) {
		pqdur := PqDuration(time.Second)
		err := pqdur.Scan([]uint8("2:50"))
		require.NoError(t, err)
	})

	t.Run("Test scanning with error", func(t *testing.T) {
		pqdur := PqDuration(time.Second)
		err := pqdur.Scan([]uint8("some unparseable string"))
		require.Error(t, err)
	})

	t.Run("Test getting value", func(t *testing.T) {
		pqdur := PqDuration(time.Hour)
		interval, err := pqdur.Value()
		require.NoError(t, err)
		require.Equal(t, "1h0m0s", interval)
	})
}
