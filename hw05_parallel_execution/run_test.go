package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tests := []struct {
			description       string
			taskCount         int
			workersCount      int
			maxErrorsCount    int
			expectedTasksDone int
		}{
			{
				description:       "50 workers, 23 max errors",
				taskCount:         50,
				workersCount:      10,
				maxErrorsCount:    23,
				expectedTasksDone: 10 + 23,
			},
			{
				description:       "50 workers, 1 max errors",
				taskCount:         50,
				workersCount:      10,
				maxErrorsCount:    1,
				expectedTasksDone: 10 + 2,
			},
			{
				description:       "50 workers, 0 errors allowed",
				taskCount:         50,
				workersCount:      5,
				maxErrorsCount:    0,
				expectedTasksDone: 5 + 2,
			},
			{
				description:       "50 workers, 0 errors allowed",
				taskCount:         50,
				workersCount:      5,
				maxErrorsCount:    -2,
				expectedTasksDone: 5 + 2,
			},
		}

		for _, tc := range tests {
			tc := tc
			t.Run(tc.description, func(t *testing.T) {
				tasks := make([]Task, 0, tc.taskCount)

				var runTasksCount int32

				for i := 0; i < tc.taskCount; i++ {
					err := fmt.Errorf("error from task %d", i)
					tasks = append(tasks, func() error {
						time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
						atomic.AddInt32(&runTasksCount, 1)
						return err
					})
				}

				err := Run(tasks, tc.workersCount, tc.maxErrorsCount)

				require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
				require.LessOrEqual(t, runTasksCount, int32(tc.expectedTasksDone), "extra tasks were started")
			})
		}
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})

	t.Run("edge case: tasks length equal to zero", func(t *testing.T) {
		var tasks []Task
		err := Run(tasks, 20, 10)
		require.Truef(t, errors.Is(err, nil), "actual err - %v", err)
	})

	t.Run("test case with require.eventually", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}
