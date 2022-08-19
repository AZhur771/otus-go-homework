package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Worker(
	taskChannel chan Task,
	errorCounter *int32,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for {
		select {
		case task, ok := <-taskChannel:
			if !ok {
				return
			}

			if taskResult := task(); taskResult != nil {
				atomic.AddInt32(errorCounter, 1)
			}
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	// early exit
	if len(tasks) == 0 {
		return nil
	}

	// if number of allowed errors is less or equal to zero, then no errors allowed
	if m <= 0 {
		m = 1
	}

	tasksLength := len(tasks)

	// if there are fewer tasks than goroutines, then use fewer goroutines
	if tasksLength < n {
		n = tasksLength
	}

	var wg sync.WaitGroup

	wg.Add(n)

	tasksCh := make(chan Task)
	var errorCounter int32
	errorCounter = 0 // count errors

	for i := 0; i < n; i++ {
		go Worker(tasksCh, &errorCounter, &wg)
	}

	for _, task := range tasks {
		// emergency exit
		if atomic.LoadInt32(&errorCounter) >= int32(m) {
			break
		}
		tasksCh <- task
	}

	close(tasksCh)
	wg.Wait()

	if errorCounter >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
