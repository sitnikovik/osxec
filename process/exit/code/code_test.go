package code_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/osxec/process/exit/code"
)

func TestParseCodess(t *testing.T) {
	t.Parallel()
	t.Run("valid uint8", func(t *testing.T) {
		t.Parallel()
		var n uint8 = 5
		got, err := code.ParseCode(n)
		assert.NoError(t, err)
		assert.Equal(t, code.Code(n), got)
	})
	t.Run("valid int", func(t *testing.T) {
		t.Parallel()
		var n int = 10
		got, err := code.ParseCode(n)
		assert.NoError(t, err)
		assert.Equal(t, code.Code(n), got)
	})
	t.Run("out of range int", func(t *testing.T) {
		t.Parallel()
		var n int = 256
		got, err := code.ParseCode(n)
		assert.ErrorIs(t, err, code.ErrNotParsable)
		assert.Equal(t, code.Failure, got)
	})
	t.Run("out of range negative int", func(t *testing.T) {
		t.Parallel()
		var n int = -1
		got, err := code.ParseCode(n)
		assert.ErrorIs(t, err, code.ErrNotParsable)
		assert.Equal(t, code.Failure, got)
	})
	t.Run("valid string", func(t *testing.T) {
		t.Parallel()
		var s string = "15"
		got, err := code.ParseCode(s)
		assert.NoError(t, err)
		assert.Equal(t, code.Code(15), got)
	})
	t.Run("word", func(t *testing.T) {
		t.Parallel()
		got, err := code.ParseCode("foo")
		assert.Equal(t, code.Failure, got)
		assert.ErrorIs(t, err, code.ErrNotParsable)
	})
	t.Run("out of range string", func(t *testing.T) {
		t.Parallel()
		got, err := code.ParseCode("300")
		assert.Equal(t, code.Failure, got)
		assert.ErrorIs(t, err, code.ErrNotParsable)
	})
	t.Run("out of range negative string", func(t *testing.T) {
		t.Parallel()
		got, err := code.ParseCode("-13")
		assert.Equal(t, code.Failure, got)
		assert.ErrorIs(t, err, code.ErrNotParsable)
	})
}

func TestCode_Int(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    code.Code
		want int
	}{
		{
			name: "success",
			c:    code.Success,
			want: 0,
		},
		{
			name: "failure",
			c:    code.Failure,
			want: 1,
		},
		{
			name: "custom 128",
			c:    code.Code(128),
			want: 128,
		},
		{
			name: "custom 1",
			c:    code.Code(1),
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Int()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCode_Succeeded(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		c    code.Code
		want bool
	}{
		{
			name: "success",
			c:    code.Success,
			want: true,
		},
		{
			name: "failure",
			c:    code.Failure,
			want: false,
		},
		{
			name: "custom 128",
			c:    code.Code(128),
			want: false,
		},
		{
			name: "custom 1",
			c:    code.Code(1),
			want: false,
		},
		{
			name: "custom zero",
			c:    code.Code(0),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.c.Succeeded()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestCode_Equals(t *testing.T) {
	t.Parallel()
	t.Run("1 from string eq invalid argument code", func(t *testing.T) {
		t.Parallel()
		c, err := code.ParseCode("1")
		require.NoError(t, err)
		assert.True(t, c.Equals(code.Failure))
	})
}
