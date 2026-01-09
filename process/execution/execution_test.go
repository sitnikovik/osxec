package execution_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sitnikovik/osxec/command/output"
	"github.com/sitnikovik/osxec/internal/test/fake"
	exec "github.com/sitnikovik/osxec/process/execution"
	"github.com/sitnikovik/osxec/process/exit/code"
)

func TestExecution_Code(t *testing.T) {
	t.Parallel()
	type fields struct {
		bb  []byte
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   code.Code
	}{
		{
			name: "ok 0",
			fields: fields{
				bb: []byte("foo"),
			},
			want: code.Success,
		},
		{
			name: "code 1 from error",
			fields: fields{
				err: errors.New("exit status 1"),
			},
			want: code.Failure,
		},
		{
			name: "code not found in error",
			fields: fields{
				err: fake.Err,
			},
			want: code.Failure,
		},
		{
			name: "code not parsed in error",
			fields: fields{
				err: errors.New("exit status foo"),
			},
			want: code.Failure,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := exec.
				NewExecution(
					tt.fields.bb,
					tt.fields.err,
				).
				Code()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExecution_Output(t *testing.T) {
	t.Parallel()
	type fields struct {
		bb  []byte
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   output.Output
	}{
		{
			name: "ok",
			fields: fields{
				bb: []byte("foo"),
			},
			want: output.NewOutput([]byte("foo")),
		},
		{
			name: "empty",
			fields: fields{
				bb: []byte(""),
			},
			want: output.NewOutput([]byte("")),
		},
		{
			name: "nil",
			want: output.NewOutput(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := exec.
				NewExecution(
					tt.fields.bb,
					tt.fields.err,
				).
				Output()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestExecution_Err(t *testing.T) {
	t.Parallel()
	type fields struct {
		bb  []byte
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "ok",
			fields: fields{
				err: fake.Err,
			},
			wantErr: fake.Err,
		},
		{
			name: "nil err",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := exec.
				NewExecution(
					tt.fields.bb,
					tt.fields.err,
				).
				Err()
			assert.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestExecution_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		bb  []byte
		err error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ok",
			fields: fields{
				bb: []byte("foo"),
			},
			want: "foo",
		},
		{
			name: "empty",
			fields: fields{
				bb: []byte(""),
			},
		},
		{
			name: "nil",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := exec.
				NewExecution(
					tt.fields.bb,
					tt.fields.err,
				).
				String()
			assert.Equal(t, tt.want, got)
		})
	}
}
