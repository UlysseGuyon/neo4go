package errors

import (
	"strings"
	"testing"
)

func TestSecurityError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *SecurityError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &SecurityError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("SecurityError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecurityError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *SecurityError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &SecurityError{Err: "A typical error"},
			want: SecurityErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &SecurityError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("SecurityError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *AuthError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &AuthError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("AuthError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *AuthError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &AuthError{Err: "A typical error"},
			want: AuthErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &AuthError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("AuthError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *ClientError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &ClientError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("ClientError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *ClientError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &ClientError{Err: "A typical error"},
			want: ClientErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &ClientError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("ClientError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransientError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *TransientError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &TransientError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("TransientError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransientError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *TransientError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &TransientError{Err: "A typical error"},
			want: TransientErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &TransientError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("TransientError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *SessionError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &SessionError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("SessionError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSessionError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *SessionError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &SessionError{Err: "A typical error"},
			want: SessionErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &SessionError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("SessionError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnavailableError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *UnavailableError
		want string
	}{
		{
			name: "Should be exactly the same error string",
			err:  &UnavailableError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("UnavailableError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnavailableError_FmtError(t *testing.T) {
	tests := []struct {
		name string
		err  *UnavailableError
		want string
	}{
		{
			name: "Should contain the error type",
			err:  &UnavailableError{Err: "A typical error"},
			want: UnavailableErrorTypeName,
		},
		{
			name: "Should contain the whole error string",
			err:  &UnavailableError{Err: "A typical error"},
			want: "A typical error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.FmtError(); !strings.Contains(got, tt.want) {
				t.Errorf("UnavailableError.FmtError() = %v, want %v", got, tt.want)
			}
		})
	}
}
