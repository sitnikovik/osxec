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

func TestString(t *testing.T) {
	t.Parallel()
	t.Run("cmd and args", func(t *testing.T) {
		t.Parallel()
		got := NewCommand("docker stop", "container_id", "--force").String()
		assert.Equal(
			t,
			"docker stop container_id --force",
			got,
		)
	})
	t.Run("only cmd", func(t *testing.T) {
		t.Parallel()
		got := NewCommand("docker stop").String()
		assert.Equal(t, "docker stop", got)
	})
	t.Run("only args", func(t *testing.T) {
		t.Parallel()
		got := NewCommand("", "container_id", "--force").String()
		assert.Empty(t, got)
	})
	t.Run("zero value args", func(t *testing.T) {
		t.Parallel()
		var (
			name string
			args []string
		)
		got := NewCommand(name, args...).String()
		assert.Empty(t, got)
	})
}

func TestCommand_Name(t *testing.T) {
	t.Parallel()
	type fields struct {
		args []string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "with name",
			fields: fields{
				name: "docker run",
				args: []string{"container_id", "--force"},
			},
			want: "docker run",
		},
		{
			name: "empty name",
			fields: fields{
				args: []string{"container_id", "--force"},
			},
		},
		{
			name: "zero value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewCommand(
				tt.fields.name,
				tt.fields.args...,
			).Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCommand_Args(t *testing.T) {
	t.Parallel()
	type fields struct {
		args []string
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "with args",
			fields: fields{
				name: "docker run",
				args: []string{"container_id", "--force"},
			},
			want: []string{"container_id", "--force"},
		},
		{
			name: "empty args",
			fields: fields{
				name: "docker run",
				args: []string{},
			},
		},
		{
			name: "zero value",
			fields: fields{
				name: "",
				args: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewCommand(
				tt.fields.name,
				tt.fields.args...,
			).Args()
			assert.Equal(t, tt.want, got)
		})
	}
}
