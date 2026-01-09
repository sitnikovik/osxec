package shell

import (
	"context"

	"github.com/sitnikovik/osxec/command"
	proc "github.com/sitnikovik/osxec/process/execution"
)

// Shell is a fake implementation of the Shell interface for testing purposes.
type Shell struct {
	// resp holds predefined responses for specific commands.
	resp map[string]proc.Execution
}

// Option defines a function type for configuring the Shell instance.
type Option func(*Shell)

// WithResponse sets a predefined response for a specific command.
//
// Parameters:
//   - cmd: The command for which the response is to be set.
//   - res: The execution result to be returned when the command is executed.
func WithResponse(
	cmd command.Command,
	res proc.Execution,
) Option {
	return func(s *Shell) {
		s.resp[cmd.String()] = res
	}
}

// NewShell creates and returns a new Shell instance to execute system commands.
func NewShell(opts ...Option) *Shell {
	sh := &Shell{
		resp: make(map[string]proc.Execution),
	}
	for _, opt := range opts {
		opt(sh)
	}
	return sh
}

// Execution executes the given command and returns its execution result.
//
// If context is done before execution, it returns an error execution with context error.
//
// Parameters:
//   - ctx: The context for managing the command execution lifecycle.
//   - cmd: The command to be executed.
func (s Shell) Execution(
	ctx context.Context,
	cmd command.Command,
) proc.Execution {
	if res, ok := s.resp[cmd.String()]; ok {
		return res
	}
	panic("response not found for command '" + cmd.String() + "'")
}
