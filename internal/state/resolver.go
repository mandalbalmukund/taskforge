package state

import (
	"fmt"
	"strings"
)

func (s *SharedState) Resolve(
	ref string,
) (interface{}, error) {

	ref = strings.TrimSpace(ref)

	ref = strings.TrimPrefix(ref, "{{")
	ref = strings.TrimSuffix(ref, "}}")

	parts := strings.Split(ref, ".")

	if len(parts) < 3 {
		return nil,
			fmt.Errorf(
				"invalid reference",
			)
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks :=
		s.data["tasks"].
			(map[string]interface{})

	taskData, ok :=
		tasks[parts[1]]

	if !ok {
		return nil,
			fmt.Errorf(
				"task output not found",
			)
	}

	output :=
		taskData.
			(map[string]interface{})

	value, ok :=
		output[parts[2]]

	if !ok {
		return nil,
			fmt.Errorf(
				"value not found",
			)
	}

	return value, nil
}