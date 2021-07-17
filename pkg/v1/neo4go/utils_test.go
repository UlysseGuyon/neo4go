package neo4go

import (
	"reflect"
	"testing"
)

func TestGetValueElem(t *testing.T) {
	type args struct {
		val reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueElem(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
