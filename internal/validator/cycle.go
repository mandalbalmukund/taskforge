package validator

import "taskforge/internal/model"

func cycleExists(tasks []model.TaskDefinition) bool {

	graph := make(map[string][]string)

	for _, task := range tasks {
		graph[task.ID] = task.DependsOn
	}

	visited := map[string]bool{}
	recStack := map[string]bool{}

	var dfs func(string) bool

	dfs = func(node string) bool {

		if recStack[node] {
			return true
		}

		if visited[node] {
			return false
		}

		visited[node] = true
		recStack[node] = true

		for _, dep := range graph[node] {
			if dfs(dep) {
				return true
			}
		}

		recStack[node] = false

		return false
	}

	for _, task := range tasks {
		if dfs(task.ID) {
			return true
		}
	}

	return false
}