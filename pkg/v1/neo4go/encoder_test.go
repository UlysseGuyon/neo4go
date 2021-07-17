package neo4go

import (
	"reflect"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	type args struct {
		opt *EncoderOptions
	}
	tests := []struct {
		name string
		args args
		want Encoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEncoder(tt.args.opt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goEncoder_Encode(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		encoder *neo4goEncoder
		args    args
		want    InputStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.encoder.Encode(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goEncoder_getDefaultHook(t *testing.T) {
	tests := []struct {
		name    string
		encoder *neo4goEncoder
		want    EncodeHookFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.encoder.getDefaultHook(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goEncoder.getDefaultHook() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComposeEncodeHookFunc(t *testing.T) {
	type args struct {
		hooks []EncodeHookFunc
	}
	tests := []struct {
		name string
		args args
		want EncodeHookFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ComposeEncodeHookFunc(tt.args.hooks...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ComposeEncodeHookFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
