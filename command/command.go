package command

import (
	"strings"
)

// Command represents a system command with its name and arguments.
type Command struct {
	// args holds the arguments for the command.
	args []string
	// name is the name of the command to be executed.
	name string
}

// NewCommand creates and returns a new Command instance with the specified name and arguments.
//
// Parameters:
//   - name: The name of the command to be executed.
//   - args: A variadic list of arguments for the command.
func NewCommand(
	name string,
	args ...string,
) Command {
	var aa []string
	if len(args) > 0 {
		aa = make([]string, len(args))
		copy(aa, args)
	}
	return Command{
		args: aa,
		name: name,
	}
}

// String returns a string representation of the Command,
// combining the command name and its arguments.
func (c Command) String() string {
	if c.name == "" {
		return ""
	}
	if len(c.args) == 0 {
		return c.name
	}
	return c.name + " " + strings.Join(c.args, " ")
}

// Name returns the name of the Command.
func (c Command) Name() string {
	return c.name
}

// Args returns a copy of the arguments slice of the Command.
func (c Command) Args() []string {
	if len(c.args) == 0 {
		return nil
	}
	res := make([]string, len(c.args))
	copy(res, c.args)
	return res
}
