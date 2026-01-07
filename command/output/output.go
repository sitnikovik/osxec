package output

import "strings"

// Output represents the output of a system command execution.
type Output struct {
	// stdout holds the standard output of the command.
	stdout []byte
}

// NewOutput creates and returns a new Output instance
// with the specified standard output.
//
// Parameters:
//   - stdout: A byte slice representing the standard output of the command.
func NewOutput(stdout []byte) Output {
	return Output{
		stdout: stdout,
	}
}

// Len returns the length of the output.
func (o Output) Len() int {
	return len(o.stdout)
}

// Empty checks if the output is empty.
func (o Output) Empty() bool {
	return o.Len() == 0
}

// String returns a string representation of the output.
func (o Output) String() string {
	return string(o.stdout)
}

// Lines splits the output into lines and returns them as a Lines instance.
//
// It splits the output string at each newline character.
func (o Output) Lines() Lines {
	if o.Empty() {
		return NewLines(nil)
	}
	return NewLines(
		strings.Split(
			o.String(),
			"\n",
		),
	)
}

// Bytes returns the output as a byte slice.
//
// It returns a copy of the internal byte slice
// to prevent external modification.
func (o Output) Bytes() []byte {
	if o.Empty() {
		return nil
	}
	res := make([]byte, len(o.stdout))
	copy(res, o.stdout)
	return res
}
