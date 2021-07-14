package neo4go

import (
	"reflect"
	"testing"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Test_newEmptyRecordMap(t *testing.T) {
	tests := []struct {
		name string
		want RecordMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEmptyRecordMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEmptyRecordMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSingle(t *testing.T) {
	type args struct {
		from QueryResult
		err  error
	}
	tests := []struct {
		name  string
		args  args
		want  *RecordMap
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Single(tt.args.from, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Single() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Single() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCollect(t *testing.T) {
	type args struct {
		from QueryResult
		err  error
	}
	tests := []struct {
		name  string
		args  args
		want  []RecordMap
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Collect(tt.args.from, tt.args.err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collect() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Collect() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_decodeItemInRecordMap(t *testing.T) {
	type args struct {
		key          string
		value        interface{}
		resultRecord *RecordMap
	}
	tests := []struct {
		name string
		args args
		want internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeItemInRecordMap(tt.args.key, tt.args.value, tt.args.resultRecord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeItemInRecordMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeMap(t *testing.T) {
	type args struct {
		mapInterface map[string]interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  *RecordMap
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := decodeMap(tt.args.mapInterface)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeMap() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("decodeMap() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_newQueryResult(t *testing.T) {
	type args struct {
		result neo4j.Result
	}
	tests := []struct {
		name string
		args args
		want QueryResult
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newQueryResult(tt.args.result); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newQueryResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_Keys(t *testing.T) {
	tests := []struct {
		name    string
		res     *queryResult
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.res.Keys()
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResult.Keys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResult.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_Next(t *testing.T) {
	tests := []struct {
		name string
		res  *queryResult
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.res.Next(); got != tt.want {
				t.Errorf("queryResult.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_Err(t *testing.T) {
	tests := []struct {
		name    string
		res     *queryResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.res.Err(); (err != nil) != tt.wantErr {
				t.Errorf("queryResult.Err() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_queryResult_Record(t *testing.T) {
	tests := []struct {
		name  string
		res   *queryResult
		want  *RecordMap
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.res.Record()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResult.Record() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("queryResult.Record() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_queryResult_Summary(t *testing.T) {
	tests := []struct {
		name    string
		res     *queryResult
		want    neo4j.ResultSummary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.res.Summary()
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResult.Summary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResult.Summary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_Consume(t *testing.T) {
	tests := []struct {
		name    string
		res     *queryResult
		want    neo4j.ResultSummary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.res.Consume()
			if (err != nil) != tt.wantErr {
				t.Errorf("queryResult.Consume() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResult.Consume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResult_RawResult(t *testing.T) {
	tests := []struct {
		name string
		res  *queryResult
		want neo4j.Result
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.res.RawResult(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queryResult.RawResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
