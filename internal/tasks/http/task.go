package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"time"

	"taskforge/internal/model"
	"taskforge/internal/registry"
)

type Task struct {
	client *nethttp.Client
}

func New() *Task {
	return &Task{
		client: &nethttp.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (t *Task) Type() string {
	return "http"
}

func (t *Task) Execute(
	ctx context.Context,
	task model.TaskDefinition,
	state map[string]interface{},
) (registry.TaskResult, error) {

	method := task.Config["method"].(string)
	url := task.Config["url"].(string)

	var bodyBytes []byte

	if body, ok := task.Config["body"]; ok {
		bodyBytes, _ = json.Marshal(body)
	}

	req, err := nethttp.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(bodyBytes),
	)

	if err != nil {
		return registry.TaskResult{}, err
	}

	resp, err := t.client.Do(req)

	if err != nil {
		return registry.TaskResult{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 &&
		resp.StatusCode < 500 {

		return registry.TaskResult{},
			fmt.Errorf("permanent failure: %d",
				resp.StatusCode)
	}

	return registry.TaskResult{
		Output: map[string]interface{}{
			"statusCode": resp.StatusCode,
		},
	}, nil
}