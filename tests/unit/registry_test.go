package unit

import (
	"testing"

	"taskforge/internal/registry"
)

type MockTask struct{}

func (m MockTask) Type() string {
	return "mock"
}

func TestRegistryRegistration(
	t *testing.T,
) {

	r := registry.New()

	r.Register(MockTask{})

	if !r.Exists("mock") {
		t.Fatal(
			"handler not registered",
		)
	}
}