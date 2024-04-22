package types

import "context"

// Executor represents an interface for executing tasks.
type Executor interface {
	// New creates a new instance of the Executor.
	New() Executor

	// ID returns the unique identifier of the Executor.
	ID() string

	// Name returns the name of the Executor.
	Name() string

	// Description returns the description of the Executor.
	Description() string

	// InputRules returns the input rules for the Executor.
	InputRules() map[string]interface{}

	// OutputRules returns the output rules for the Executor.
	OutputRules() map[string]interface{}

	// Validate validates the task and its dependencies.
	Validate(ctx context.Context, task *Task, tasks []*Task) error

	// Execute executes the task and its dependencies.
	Execute(ctx context.Context, task *Task, tasks []*Task) (interface{}, error)
}

type Task struct {
	ID             string           `json:"id" yaml:"id" validate:"required"`
	Name           string           `json:"name" yaml:"name"`
	Description    string           `json:"description" yaml:"description"`
	Executor       string           `json:"executor" yaml:"executor" validate:"required"`
	DependsOn      []string         `json:"dependsOn" yaml:"dependsOn"`
	GlobalInput    interface{}      `json:"globalInput,omitempty" yaml:"globalInput,omitempty"`
	Input          interface{}      `json:"input,omitempty" yaml:"input,omitempty" validate:"required"`
	Output         interface{}      `json:"output,omitempty" yaml:"output,omitempty"`
	Condition      interface{}      `json:"condition,omitempty" yaml:"condition,omitempty"`
	Duration       int64            `json:"duration,omitempty" yaml:"duration,omitempty"`
	LatestDuration int64            `json:"latestDuration,omitempty" yaml:"latestDuration,omitempty"`
	IsExecuting    bool             `yaml:"isExecuting" json:"isExecuting"`
	Succeeded      bool             `yaml:"succeeded" json:"succeeded"`
	NextTaskId     string           `json:"nextTaskId,omitempty" yaml:"nextTaskId,omitempty"`
	Error          error            `json:"error,omitempty" yaml:"error,omitempty"`
	Dependencies   map[string]*Task `json:"-" yaml:"-"`
}
