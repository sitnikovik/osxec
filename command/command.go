package command

import "strings"

// Command represents a system command with its name and arguments.
type Command struct {
	// args holds the arguments for the command.
	args []string
	// name is the name of the command to be executed.
	name string
}

// NewCommand creates and returns a new Command instance with the specified.
//
// Parameters:
//   - name: The name of the command to be executed.
//   - args: A variadic list of arguments for the command.
func NewCommand(
	name string,
	args ...string,
) Command {
	return Command{
		args: args,
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
