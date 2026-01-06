package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	t.Parallel()
	t.Run("cmd and args", func(t *testing.T) {
		t.Parallel()
		assert.Equal(
			t,
			Command{
				args: []string{"container_id", "--force"},
				name: "docker stop",
			},
			NewCommand(
				"docker stop",
				"container_id", "--force",
			),
		)
	})
	t.Run("only cmd", func(t *testing.T) {
		t.Parallel()
		got := NewCommand("docker stop")
		want := Command{
			name: "docker stop",
		}
		assert.Equal(t, want, got)
		assert.Nil(t, got.args)
	})
	t.Run("only args", func(t *testing.T) {
		t.Parallel()
		var (
			name string
			args []string = []string{"container_id", "--force"}
			want          = Command{
				args: []string{"container_id", "--force"},
			}
		)
		got := NewCommand(name, args...)
		assert.Equal(t, want, got)
	})
	t.Run("zero value args", func(t *testing.T) {
		t.Parallel()
		var (
			name string
			args []string
			want Command
		)
		got := NewCommand(name, args...)
		assert.Equal(t, want, got)
	})
}
