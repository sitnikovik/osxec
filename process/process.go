package process

import (
	"context"

	"github.com/sitnikovik/osxec/command"
	exec "github.com/sitnikovik/osxec/process/execution"
)

// Shell defines an interface for executing commands in a shell environment.
type Shell interface {
	// Execution executes the given command and returns its execution result.
	//
	// Parameters:
	//   - ctx: The context for managing the command execution lifecycle.
	//   - cmd: The command to be executed.
	Execution(
		ctx context.Context,
		cmd command.Command,
	) exec.Execution
}

// Process represents a system process that can execute commands using a shell.
type Process struct {
	// shell is the shell environment used to execute commands.
	shell Shell
	// cmd is the command to be executed by the process.
	cmd command.Command
}

// NewProcess creates and returns a new Process instance with the specified shell and command.
//
// Parameters:
//   - shell: The shell environment used to execute commands.
//   - cmd: The command to be executed by the process.
func NewProcess(
	shell Shell,
	cmd command.Command,
) Process {
	return Process{
		shell: shell,
		cmd:   cmd,
	}
}

// Execution returns the execution result of the process
// by executing its command in the specified shell.
func (p Process) Execution(ctx context.Context) exec.Execution {
	return p.shell.Execution(ctx, p.cmd)
}
