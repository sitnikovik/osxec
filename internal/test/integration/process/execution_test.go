//go:build integration
// +build integration

package process_test

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

func TestExecution(t *testing.T) {
	t.Parallel()
	t.Run("echo hello world", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		res := process.
			NewProcess(
				shell.NewShell(),
				command.NewCommand("echo", "Hello, World!"),
			).
			Execution(ctx)
		require.NoError(t, res.Err())
		assert.Equal(t, "Hello, World!\n", res.String())
	})
	t.Run("wrong command", func(t *testing.T) {
		t.Parallel()
		ctx := context.Background()
		res := process.
			NewProcess(
				shell.NewShell(),
				command.NewCommand("git", "sdasd"),
			).
			Execution(ctx)
		assert.Error(t, res.Err())
	})
	t.Run("context done", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(
			context.Background(),
			1*time.Second,
		)
		defer cancel()
		res := process.
			NewProcess(
				shell.NewShell(),
				command.NewCommand("sleep", "3"),
			).
			Execution(ctx)
		time.Sleep(2 * time.Second)
		err := res.Err()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "signal: killed")
		assert.ErrorIs(t, err, context.DeadlineExceeded)
	})
	t.Run("context canceled", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithCancel(
			context.Background(),
		)
		cancel()
		res := process.
			NewProcess(
				shell.NewShell(),
				command.NewCommand("echo", "Hello"),
			).
			Execution(ctx)
		err := res.Err()
		require.Error(t, err)
		assert.ErrorIs(t, err, context.Canceled)
	})
	t.Run("exit 128", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(
			context.Background(),
			1*time.Second,
		)
		defer cancel()
		res := process.NewProcess(
			shell.NewShell(),
			command.NewCommand("sh", "-c", "exit 128"),
		).Execution(ctx)
		require.Error(t, res.Err())
		assert.Equal(t, 128, res.Code().Int())
	})
}
