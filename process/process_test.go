package process_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/osxec/command"
	fakeShell "github.com/sitnikovik/osxec/internal/test/fake/shell"
	"github.com/sitnikovik/osxec/process"
	exec "github.com/sitnikovik/osxec/process/execution"
)

func TestProcess_Execution(t *testing.T) {
	t.Parallel()
	type fields struct {
		shell process.Shell
		cmd   command.Command
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   exec.Execution
	}{
		{
			name: "ok",
			fields: fields{
				shell: fakeShell.NewShell(
					fakeShell.WithResponse(
						command.NewCommand("ping"),
						exec.NewExecution([]byte("pong\n"), nil),
					),
				),
				cmd: command.NewCommand("ping"),
			},
			args: args{
				ctx: context.Background(),
			},
			want: exec.NewExecution([]byte("pong\n"), nil),
		},
		{
			name: "command error",
			fields: fields{
				shell: fakeShell.NewShell(
					fakeShell.WithResponse(
						command.NewCommand("docker", "asdsaofjdk"),
						exec.NewExecution(
							[]byte("docker: 'asdsaofjdk' is not a docker command.\n"),
							errors.New("exit status 1"),
						),
					),
				),
				cmd: command.NewCommand("docker", "asdsaofjdk"),
			},
			args: args{
				ctx: context.Background(),
			},
			want: exec.NewExecution(
				[]byte("docker: 'asdsaofjdk' is not a docker command.\n"),
				errors.New("exit status 1"),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := process.
				NewProcess(
					tt.fields.shell,
					tt.fields.cmd,
				).
				Execution(tt.args.ctx)
			assert.Equal(t, tt.want, got)
		})
	}
}
