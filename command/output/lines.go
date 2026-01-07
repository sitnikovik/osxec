package output

// Lines represents a collection of output lines from a command execution.
//
// It is essentially a slice of strings, where each string is a line of output.
type Lines []string

// NewLines creates and returns a new Lines instance
// from the provided slice of strings.
//
// Parameters:
//   - lines: A slice of strings representing the lines of output.
func NewLines(lines []string) Lines {
	return Lines(lines)
}

// Len returns the number of lines.
func (l Lines) Len() int {
	return len(l)
}

// Empty checks if there are no lines.
func (l Lines) Empty() bool {
	return l.Len() == 0
}

// First returns the first line.
//
// It returns an empty string if there are no lines.
func (l Lines) First() string {
	if l.Empty() {
		return ""
	}
	return l[0]
}

// Last returns the last line.
//
// It returns an empty string if there are no lines.
func (l Lines) Last() string {
	if l.Empty() {
		return ""
	}
	return l[l.Len()-1]
}

// At returns the line at the specified index.
//
// If idx is negative, it is interpreted as an offset from the end of the
// slice (e.g., -1 refers to the last line, -2 to the second-to-last, etc.).
// It returns an empty string if there are no lines or if the index is out of
// bounds.
func (l Lines) At(idx int) string {
	n := l.Len()
	if idx < 0 {
		idx = n + idx
	}
	if idx < 0 || idx >= n {
		return ""
	}
	return l[idx]
}
