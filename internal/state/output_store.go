package state

func (s *SharedState) StoreTaskOutput(
	taskID string,
	output map[string]interface{},
) {

	s.mu.Lock()
	defer s.mu.Unlock()

	tasks :=
		s.data["tasks"].
			(map[string]interface{})

	tasks[taskID] = output
}