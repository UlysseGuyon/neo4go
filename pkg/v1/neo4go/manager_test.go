package neo4go

import (
	"reflect"
	"testing"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
)

func TestNewManager(t *testing.T) {
	type args struct {
		options ManagerOptions
	}
	tests := []struct {
		name  string
		args  args
		want  Manager
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewManager(tt.args.options)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewManager() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewManager() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_manager_Init(t *testing.T) {
	type args struct {
		options ManagerOptions
	}
	tests := []struct {
		name string
		m    *manager
		args args
		want internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.init(tt.args.options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_IsConnected(t *testing.T) {
	tests := []struct {
		name string
		m    *manager
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.IsConnected(); got != tt.want {
				t.Errorf("manager.IsConnected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_Close(t *testing.T) {
	tests := []struct {
		name string
		m    *manager
		want internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Close(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_manager_Query(t *testing.T) {
	type args struct {
		queryParams QueryParams
	}
	tests := []struct {
		name  string
		m     *manager
		args  args
		want  QueryResult
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.m.Query(tt.args.queryParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Query() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("manager.Query() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_manager_Transaction(t *testing.T) {
	type args struct {
		transactionGlobalParams TransactionParams
	}
	tests := []struct {
		name  string
		m     *manager
		args  args
		want  QueryResult
		want1 internalErr.Neo4GoError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.m.Transaction(tt.args.transactionGlobalParams)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("manager.Transaction() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("manager.Transaction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
