package output

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLines_Len(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "ok",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			want: 3,
		},
		{
			name:  "empty",
			lines: []string{},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewLines(tt.lines).Len()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLines_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  bool
	}{
		{
			name: "ok",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
		},
		{
			name:  "empty",
			lines: []string{},
			want:  true,
		},
		{
			name: "nil",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewLines(tt.lines).Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestLines_First(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{
			name: "ok",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			want: "line1",
		},
		{
			name:  "empty",
			lines: []string{},
		},
		{
			name:  "nil",
			lines: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewLines(tt.lines).First()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLines_Last(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		want  string
	}{
		{
			name: "ok",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			want: "line3",
		},
		{
			name:  "empty",
			lines: []string{},
		},
		{
			name:  "nil",
			lines: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewLines(tt.lines).Last()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLines_At(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		lines []string
		idx   int
		want  string
	}{
		{
			name: "index 0",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			idx:  0,
			want: "line1",
		},
		{
			name: "index 1",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			idx:  1,
			want: "line2",
		},
		{
			name: "index 2",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			idx:  2,
			want: "line3",
		},
		{
			name: "index -1",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			idx:  -1,
			want: "line3",
		},
		{
			name: "index -2",
			lines: []string{
				"line1",
				"line2",
				"line3",
			},
			idx:  -2,
			want: "line2",
		},
		{
			name:  "empty",
			lines: []string{},
		},
		{
			name:  "nil",
			lines: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := NewLines(tt.lines).At(tt.idx)
			assert.Equal(t, tt.want, got)
		})
	}
}
