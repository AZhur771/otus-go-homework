package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	tmpCh := in

	// early exit
	if len(stages) == 0 {
		return tmpCh
	}

	// cancel stage
	cancellator := func(done In, in In) Out {
		out := make(Bi)
		go func() {
			defer close(out)
			// active listening
			for {
				// prioritize cancellation condition
				select {
				case <-done:
					return
				default:
				}

				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					out <- v
				default:
				}
			}
		}()
		return out
	}

	for _, stage := range stages {
		tmpCh = cancellator(done, stage(tmpCh))
	}

	return tmpCh
}
