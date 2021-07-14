package errors

import (
	"errors"
	"reflect"
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

func TestToDriverError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want Neo4GoError
	}{
		{
			name: "Should return unknwon error if it's not a neo4j-go-driver error",
			args: args{
				err: errors.New("A typical error"),
			},
			want: &UnknownError{Err: "A typical error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToDriverError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDriverError() = %v, want %v", got, tt.want)
			}
		})
	}
}
