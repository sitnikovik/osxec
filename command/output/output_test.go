package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutput_Len(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		stdout []byte
		want   int
	}{
		{
			name: "with new lines",
			stdout: []byte(`line1
line2
line3`),
			want: 17,
		},
		{
			name:   "chinese",
			stdout: []byte("你好，世界"),
			want:   15,
		},
		{
			name:   "spaces",
			stdout: []byte("     "),
			want:   5,
		},
		{
			name:   "empty",
			stdout: []byte{},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewOutput(tt.stdout).Len()
			assert.Equal(t, tt.want, got)
		})
	}
}
func TestOutput_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		stdout []byte
		want   bool
	}{
		{
			name: "ok",
			stdout: []byte(`line1
line2
line3`),
		},
		{
			name:   "spaces",
			stdout: []byte("     "),
		},
		{
			name:   "empty",
			stdout: []byte{},
			want:   true,
		},
		{
			name: "nil",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewOutput(tt.stdout).Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestOutput_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		stdout []byte
		want   string
	}{
		{
			name: "with new lines",
			stdout: []byte(`line1
line2
line3`),
			want: `line1
line2
line3`,
		},
		{
			name:   "chinese",
			stdout: []byte("你好，世界"),
			want:   "你好，世界",
		},
		{
			name:   "spaces",
			stdout: []byte("     "),
			want:   "     ",
		},
		{
			name:   "empty",
			stdout: []byte{},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewOutput(tt.stdout).String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOutput_Lines(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		stdout []byte
		want   Lines
	}{
		{
			name:   "with new lines",
			stdout: []byte("line1\nline2\n"),
			want: NewLines([]string{
				"line1",
				"line2",
				"",
			}),
		},
		{
			name:   "spaces",
			stdout: []byte("     "),
			want:   NewLines([]string{"     "}),
		},
		{
			name:   "empty",
			stdout: []byte{},
			want:   NewLines(nil),
		},
		{
			name: "nil",
			want: NewLines(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewOutput(tt.stdout).Lines()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestOutput_Bytes(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		stdout []byte
		want   []byte
	}{
		{
			name:   "ok",
			stdout: []byte(`foo`),
			want:   []byte(`foo`),
		},
		{
			name:   "empty",
			stdout: []byte{},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewOutput(tt.stdout).Bytes()
			assert.Equal(t, tt.want, got)
			if len(got) > 0 && len(tt.want) > 0 {
				assert.False(t, &tt.stdout[0] == &got[0], "must be different slices")
			}
		})
	}
}
