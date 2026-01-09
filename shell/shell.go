package shell

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/sitnikovik/osxec/command"
	proc "github.com/sitnikovik/osxec/process/execution"
)

// Shell provides an implementation for executing system commands.
type Shell struct{}

// NewShell creates and returns a new Shell instance to execute system commands.
func NewShell() Shell {
	return Shell{}
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
	bb, err := exec.
		CommandContext(
			ctx,
			cmd.Name(),
			cmd.Args()...,
		).
		CombinedOutput()
	if err != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
		err = fmt.Errorf("%w: %w", ctx.Err(), err)
	}
	return proc.NewExecution(bb, err)
}
