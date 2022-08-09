package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func Worker(
	taskChannel chan Task,
	taskDoneSignal chan struct{},
	errorSignal chan struct{},
	abortSignal chan struct{},
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for {
		// prioritize abortSignal
		select {
		case <-abortSignal:
			return
		default:
		}

		select {
		case task, ok := <-taskChannel:
			if !ok {
				return
			}

			if taskResult := task(); taskResult != nil {
				errorSignal <- struct{}{}
			} else {
				taskDoneSignal <- struct{}{}
			}
		default:
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

	tasksCh := make(chan Task, tasksLength)
	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	errorCounter := 0 // count errors
	errorSignal := make(chan struct{}, tasksLength)
	defer close(errorSignal)

	taskDoneCounter := 0 // count tasks done
	taskDoneSignal := make(chan struct{}, tasksLength)
	defer close(taskDoneSignal)

	abortCh := make(chan struct{})

	for i := 0; i < n; i++ {
		go Worker(tasksCh, taskDoneSignal, errorSignal, abortCh, &wg)
	}

	for {
		select {
		case <-taskDoneSignal:
			taskDoneCounter++
			// first exit condition
			if taskDoneCounter == tasksLength {
				wg.Wait() // allow goroutines finish started tasks
				return nil
			}
		case <-errorSignal:
			errorCounter++
			// second exit condition
			if errorCounter >= m {
				close(abortCh)
				wg.Wait() // allow goroutines finish started tasks
				return ErrErrorsLimitExceeded
			}
		}
	}
}
