package neo4go

import (
	"errors"
	"reflect"
	"testing"

	internalErr "github.com/UlysseGuyon/neo4go/internal/errors"
)

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
			want: &internalErr.UnknownError{Err: "A typical error"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toDriverError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToDriverError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsInitError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Init error",
			args: args{
				err: &internalErr.InitError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInitError(tt.args.err); got != tt.want {
				t.Errorf("IsInitError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTypeError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Type error",
			args: args{
				err: &internalErr.TypeError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTypeError(tt.args.err); got != tt.want {
				t.Errorf("IsTypeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDecodingError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Decoding error",
			args: args{
				err: &internalErr.DecodingError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDecodingError(tt.args.err); got != tt.want {
				t.Errorf("IsDecodingError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsQueryError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Query error",
			args: args{
				err: &internalErr.QueryError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsQueryError(tt.args.err); got != tt.want {
				t.Errorf("IsQueryError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTransactionError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Transaction error",
			args: args{
				err: &internalErr.TransactionError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTransactionError(tt.args.err); got != tt.want {
				t.Errorf("IsTransactionError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUnknownError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Unknown error",
			args: args{
				err: &internalErr.UnknownError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUnknownError(tt.args.err); got != tt.want {
				t.Errorf("IsUnknownError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSecurityError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Security error",
			args: args{
				err: &internalErr.SecurityError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSecurityError(tt.args.err); got != tt.want {
				t.Errorf("IsSecurityError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAuthError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Auth error",
			args: args{
				err: &internalErr.AuthError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAuthError(tt.args.err); got != tt.want {
				t.Errorf("IsAuthError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsClientError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Client error",
			args: args{
				err: &internalErr.ClientError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsClientError(tt.args.err); got != tt.want {
				t.Errorf("IsClientError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsTransientError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Transient error",
			args: args{
				err: &internalErr.TransientError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTransientError(tt.args.err); got != tt.want {
				t.Errorf("IsTransientError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSessionError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Session error",
			args: args{
				err: &internalErr.SessionError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSessionError(tt.args.err); got != tt.want {
				t.Errorf("IsSessionError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsServiceUnavailableError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Should detect Service Unavailable error",
			args: args{
				err: &internalErr.UnavailableError{},
			},
			want: true,
		},
		{
			name: "Should not detect basic error",
			args: args{
				err: errors.New(""),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsServiceUnavailableError(tt.args.err); got != tt.want {
				t.Errorf("IsServiceUnavailableError() = %v, want %v", got, tt.want)
			}
		})
	}
}
