// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package v0alpha1

// +k8s:openapi-gen=true
type StatusOperatorState struct {
	// lastEvaluation is the ResourceVersion last evaluated
	LastEvaluation string `json:"lastEvaluation"`
	// state describes the state of the lastEvaluation.
	// It is limited to three possible states for machine evaluation.
	State StatusOperatorStateState `json:"state"`
	// descriptiveState is an optional more descriptive state field which has no requirements on format
	DescriptiveState *string `json:"descriptiveState,omitempty"`
	// details contains any extra information that is operator-specific
	Details map[string]interface{} `json:"details,omitempty"`
}

// NewStatusOperatorState creates a new StatusOperatorState object.
func NewStatusOperatorState() *StatusOperatorState {
	return &StatusOperatorState{}
}

// +k8s:openapi-gen=true
type Status struct {
	// operatorStates is a map of operator ID to operator state evaluations.
	// Any operator which consumes this kind SHOULD add its state evaluation information to this field.
	OperatorStates map[string]StatusOperatorState `json:"operatorStates,omitempty"`
	// additionalFields is reserved for future use
	AdditionalFields map[string]interface{} `json:"additionalFields,omitempty"`
}

// NewStatus creates a new Status object.
func NewStatus() *Status {
	return &Status{}
}

// +k8s:openapi-gen=true
type StatusOperatorStateState string

const (
	StatusOperatorStateStateSuccess    StatusOperatorStateState = "success"
	StatusOperatorStateStateInProgress StatusOperatorStateState = "in_progress"
	StatusOperatorStateStateFailed     StatusOperatorStateState = "failed"
)
