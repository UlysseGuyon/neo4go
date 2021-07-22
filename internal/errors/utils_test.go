package errors

import (
	"strings"
	"testing"
)

func Test_errorFmt(t *testing.T) {
	type args struct {
		t   string
		err string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Should contain provided type",
			args: args{
				t:   "Query",
				err: "A typical error",
			},
			want: "Query",
		},
		{
			name: "Should contain provided error",
			args: args{
				t:   "Query",
				err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := errorFmt(tt.args.t, tt.args.err); !strings.Contains(got, tt.want) {
				t.Errorf("errorFmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
