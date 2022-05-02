package neo4go

import (
	"reflect"
	"testing"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/mitchellh/mapstructure"
)

func TestDecoder(t *testing.T) {
	type args struct {
		options *mapstructure.DecoderConfig
	}
	tests := []struct {
		name string
		args args
		want Decoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDecoder(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDecoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goDecoder_decodeSingleValue(t *testing.T) {
	type args struct {
		mapInput map[string]interface{}
		output   interface{}
	}
	tests := []struct {
		name    string
		decoder *neo4goDecoder
		args    args
		want    internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.decoder.decodeSingleValue(tt.args.mapInput, tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goDecoder.decodeSingleValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goDecoder_DecodeNode(t *testing.T) {
	type args struct {
		node   interface{}
		output interface{}
	}
	tests := []struct {
		name    string
		decoder *neo4goDecoder
		args    args
		want    internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.decoder.DecodeNode(tt.args.node, tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goDecoder.DecodeNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goDecoder_DecodeRelationship(t *testing.T) {
	type args struct {
		relationship interface{}
		output       interface{}
	}
	tests := []struct {
		name    string
		decoder *neo4goDecoder
		args    args
		want    internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.decoder.DecodeRelationship(tt.args.relationship, tt.args.output); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goDecoder.DecodeRelationship() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_neo4goDecoder_DecodePath(t *testing.T) {
	type args struct {
		path                interface{}
		outputNodes         interface{}
		outputRelationships interface{}
	}
	tests := []struct {
		name    string
		decoder *neo4goDecoder
		args    args
		want    internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.decoder.DecodePath(tt.args.path, tt.args.outputNodes, tt.args.outputRelationships); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neo4goDecoder.DecodePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
