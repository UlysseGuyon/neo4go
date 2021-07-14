package errors

import (
	"strings"
	"testing"
)

func TestInitError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *InitError
		want string
	}{
		{
			name: "Should contain raw error string",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "A typical error",
		},
		{
			name: "Should contain database URI",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "bolt://localhost:7687",
		},
		{
			name: "Should contain database name",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "neo4j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("InitError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *InitError
		want string
	}{
		{
			name: "Should contain error type name",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: InitErrorTypeName,
		},
		{
			name: "Should contain raw error string",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "A typical error",
		},
		{
			name: "Should contain database URI",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "bolt://localhost:7687",
		},
		{
			name: "Should contain database name",
			err: &InitError{
				Err:    "A typical error",
				URI:    "bolt://localhost:7687",
				DBName: "neo4j",
			},
			want: "neo4j",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("InitError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *TypeError
		want string
	}{
		{
			name: "Should contain raw error string",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "A typical error",
		},
		{
			name: "Should contain first expected type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "Node",
		},
		{
			name: "Should contain second expected type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "*Node",
		},
		{
			name: "Should contain effective type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "Relationship",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("TypeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTypeError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *TypeError
		want string
	}{
		{
			name: "Should contain error type name",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: TypeErrorTypeName,
		},
		{
			name: "Should contain raw error string",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "A typical error",
		},
		{
			name: "Should contain first expected type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "Node",
		},
		{
			name: "Should contain second expected type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "*Node",
		},
		{
			name: "Should contain effective type",
			err: &TypeError{
				Err:           "A typical error",
				ExpectedTypes: []string{"Node", "*Node"},
				GotType:       "Relationship",
			},
			want: "Relationship",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("TypeError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodingError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *DecodingError
		want string
	}{
		{
			name: "Should contain raw error string",
			err: &DecodingError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("DecodingError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecodingError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *DecodingError
		want string
	}{
		{
			name: "Should contain error type name",
			err: &DecodingError{
				Err: "A typical error",
			},
			want: DecodingErrorTypeName,
		},
		{
			name: "Should contain raw error string",
			err: &DecodingError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("DecodingError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *QueryError
		want string
	}{
		{
			name: "Should contain raw error string",
			err: &QueryError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("QueryError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *QueryError
		want string
	}{
		{
			name: "Should contain error type name",
			err: &QueryError{
				Err: "A typical error",
			},
			want: QueryErrorTypeName,
		},
		{
			name: "Should contain raw error string",
			err: &QueryError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("QueryError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnknownError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *UnknownError
		want string
	}{
		{
			name: "Should contain raw error string",
			err: &UnknownError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); !strings.Contains(got, tt.want) {
				t.Errorf("UnknownError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnknownError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *UnknownError
		want string
	}{
		{
			name: "Should contain error type name",
			err: &UnknownError{
				Err: "A typical error",
			},
			want: UnknownErrorTypeName,
		},
		{
			name: "Should contain raw error string",
			err: &UnknownError{
				Err: "A typical error",
			},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("UnknownError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}
