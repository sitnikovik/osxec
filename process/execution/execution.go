package execution

import (
	"strings"

	"github.com/sitnikovik/osxec/command/output"
	"github.com/sitnikovik/osxec/process/exit/code"
)

// Execution represents the result of executing a process.
type Execution struct {
	// bb is the byte slice output of the execution.
	bb []byte
	// err is the error that occurred during execution.
	err error
}

// NewExecution creates a new Execution instance with the provided output bytes and error.
//
// Parameters:
//   - bb: A byte slice representing the output of the execution.
//   - err: An error representing any error that occurred during execution.
func NewExecution(
	bb []byte,
	err error,
) Execution {
	var bc []byte
	if len(bb) > 0 {
		bc = make([]byte, len(bb))
		copy(bc, bb)
	}
	return Execution{
		bb:  bc,
		err: err,
	}
}

// String returns the string representation of the execution output.
func (e Execution) String() string {
	return e.Output().String()
}

// Err returns the error that occurred during execution.
func (e Execution) Err() error {
	return e.err
}

// Output returns the output of the execution as an Output instance.
func (e Execution) Output() output.Output {
	return output.NewOutput(e.bb)
}

// Code returns the exit code of the execution.
func (e Execution) Code() code.Code {
	if !e.Failed() {
		return code.Success
	}
	splitted := strings.Split(
		e.err.Error(),
		"exit status ",
	)
	if len(splitted) == 2 {
		c, err := code.ParseCode(splitted[1])
		if err != nil {
			return code.Failure
		}
		return c
	}
	return code.Failure
}

// Failed indicates whether the execution failed based on the presence of an error.
func (e Execution) Failed() bool {
	return e.err != nil
}
