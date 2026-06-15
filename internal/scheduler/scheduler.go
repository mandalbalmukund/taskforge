package scheduler

import (
	"context"
	"sync"

	"taskforge/internal/engine"
	"taskforge/internal/model"
	"taskforge/internal/registry"
	"taskforge/internal/state"
)

type Scheduler struct {
	registry *registry.Registry
	workers  int
}

func New(
	reg *registry.Registry,
	workers int,
) *Scheduler {

	return &Scheduler{
		registry: reg,
		workers:  workers,
	}
}

func (s *Scheduler) Run(
	ctx context.Context,
	graph *engine.Graph,
	shared *state.SharedState,
) error {

	queue := engine.NewReadyQueue()

	for taskID, count := range graph.Dependencies {

		if count == 0 {
			queue.Push(taskID)
		}
	}

	var wg sync.WaitGroup

	for i := 0; i < s.workers; i++ {

		wg.Add(1)

		go func() {

			defer wg.Done()

			for {

				taskID, ok := queue.Pop()

				if !ok {
					return
				}

				task := graph.Tasks[taskID]

				handler, err :=
					s.registry.Get(task.Type)

				if err != nil {
					continue
				}

				result, err :=
					handler.Execute(
						ctx,
						task,
						shared.Snapshot(),
					)

				if err == nil {

					shared.StoreTaskOutput(
						taskID,
						result.Output,
					)

					for _, child :=
						range graph.Children[taskID] {

						graph.Dependencies[child]--

						if graph.Dependencies[child] == 0 {
							queue.Push(child)
						}
					}
				}
			}
		}()
	}

	wg.Wait()

	return nil
}


taskCtx, cancel :=
	timeout.WithTimeout(
		ctx,
		task.Timeout,
	)

defer cancel()

result, err :=
	handler.Execute(
		taskCtx,
		task,
		shared.Snapshot(),
	)