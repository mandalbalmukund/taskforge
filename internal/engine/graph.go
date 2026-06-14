package engine

import (
	"taskforge/internal/model"
)

type Graph struct {
	Tasks          map[string]model.TaskDefinition
	Children       map[string][]string
	Dependencies   map[string]int
}

func BuildGraph(tasks []model.TaskDefinition) *Graph {

	g := &Graph{
		Tasks:        make(map[string]model.TaskDefinition),
		Children:     make(map[string][]string),
		Dependencies: make(map[string]int),
	}

	for _, task := range tasks {

		g.Tasks[task.ID] = task

		if _, ok := g.Dependencies[task.ID]; !ok {
			g.Dependencies[task.ID] = 0
		}

		for _, dep := range task.DependsOn {

			g.Children[dep] = append(
				g.Children[dep],
				task.ID,
			)

			g.Dependencies[task.ID]++
		}
	}

	return g
}