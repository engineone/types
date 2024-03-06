package types

type Task struct {
	ID             string           `json:"id" yaml:"id" validate:"required"`
	Name           string           `json:"name" yaml:"name"`
	Description    string           `json:"description" yaml:"description"`
	Executor       string           `json:"executor" yaml:"executor" validate:"required"`
	DependsOn      []string         `json:"dependsOn" yaml:"dependsOn"`
	GlobalInput    interface{}      `json:"globalInput,omitempty" yaml:"globalInput,omitempty"`
	Input          interface{}      `json:"input,omitempty" yaml:"input,omitempty" validate:"required"`
	Output         interface{}      `json:"output,omitempty" yaml:"output,omitempty"`
	Duration       int64            `json:"duration,omitempty" yaml:"duration,omitempty"`
	LatestDuration int64            `json:"latestDuration,omitempty" yaml:"latestDuration,omitempty"`
	IsExecuting    bool             `yaml:"isExecuting" json:"isExecuting"`
	Succeeded      bool             `yaml:"succeeded" json:"succeeded"`
	NextTaskId     string           `json:"nextTaskId,omitempty" yaml:"nextTaskId,omitempty"`
	Error          error            `json:"error,omitempty" yaml:"error,omitempty"`
	Dependencies   map[string]*Task `json:"-" yaml:"-"`
	// Condition      func(input TaskOutput) bool `json:"-" yaml:"-"`
}
