package tests

import (
	"testing"

	"taskforge/internal/state"
)

func TestReferenceResolution(
	t *testing.T,
) {

	s := state.New()

	s.StoreTaskOutput(
		"build",
		map[string]interface{}{
			"imageId": "abc123",
		},
	)

	val, err :=
		s.Resolve(
			"{{tasks.build.imageId}}",
		)

	if err != nil {
		t.Fatal(err)
	}

	if val != "abc123" {
		t.Fatal(
			"unexpected value",
		)
	}
}