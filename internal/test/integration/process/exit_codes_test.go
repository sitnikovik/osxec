//go:build integration
// +build integration

package shell

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/osxec/command"
	"github.com/sitnikovik/osxec/process"
	"github.com/sitnikovik/osxec/shell"
)

func TestExitCodes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		cmd     command.Command
		code    int
		wantErr bool
	}{
		{
			name:    "success",
			cmd:     command.NewCommand("sh", "-c", "exit 0"),
			code:    0,
			wantErr: false,
		},
		{
			name:    "failure",
			cmd:     command.NewCommand("sh", "-c", "exit 1"),
			code:    1,
			wantErr: true,
		},
		{
			name:    "misuse",
			cmd:     command.NewCommand("sh", "-c", "exit 2"),
			code:    2,
			wantErr: true,
		},
		{
			name:    "not executed",
			cmd:     command.NewCommand("sh", "-c", "exit 126"),
			code:    126,
			wantErr: true,
		},
		{
			name:    "not found",
			cmd:     command.NewCommand("sh", "-c", "exit 127"),
			code:    127,
			wantErr: true,
		},
		{
			name:    "invalid argument",
			cmd:     command.NewCommand("sh", "-c", "exit 128"),
			code:    128,
			wantErr: true,
		},
		{
			name:    "terminated",
			cmd:     command.NewCommand("sh", "-c", "exit 130"),
			code:    130,
			wantErr: true,
		},
		{
			name:    "custom 45",
			cmd:     command.NewCommand("sh", "-c", "exit 45"),
			code:    45,
			wantErr: true,
		},
		{
			name:    "custom 255",
			cmd:     command.NewCommand("sh", "-c", "exit 255"),
			code:    255,
			wantErr: true,
		},
		{
			name:    "custom -123",
			cmd:     command.NewCommand("sh", "-c", "exit -123"),
			code:    133,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx, cancel := context.WithTimeout(
				context.Background(),
				1*time.Second,
			)
			defer cancel()
			res := process.
				NewProcess(
					shell.NewShell(),
					tt.cmd,
				).
				Execution(ctx)
			if tt.wantErr {
				require.Error(t, res.Err())
				assert.False(t, res.Code().Succeeded())
			} else {
				require.NoError(t, res.Err())
				assert.True(t, res.Code().Succeeded())
			}
			assert.Equal(t, tt.code, res.Code().Int())
		})
	}
}
