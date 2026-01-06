package command

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
