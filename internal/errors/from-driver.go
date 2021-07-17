package errors

// The strings representing each type of neo4go error taken from the neo4j-go-driver
const (
	SecurityErrorTypeName    = "Security"
	AuthErrorTypeName        = "Auth"
	ClientErrorTypeName      = "Client"
	TransientErrorTypeName   = "Transient"
	SessionErrorTypeName     = "Session"
	UnavailableErrorTypeName = "Unavailable Service"
)

/* ----- SECURITY ERROR ----- */

// SecurityError represents a neo4j-go-driver security error
type SecurityError struct {
	Err string
}

// Error returns the raw error string
func (err *SecurityError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *SecurityError) FmtError() string {
	return errorFmt(SecurityErrorTypeName, err.Err)
}

/* ----- AUTHENTICATION ERROR ----- */

// AuthError represents a neo4j-go-driver security error
type AuthError struct {
	Err string
}

// Error returns the raw error string
func (err *AuthError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *AuthError) FmtError() string {
	return errorFmt(AuthErrorTypeName, err.Err)
}

/* ----- CLIENT ERROR ----- */

// ClientError represents a neo4j-go-driver security error
type ClientError struct {
	Err string
}

// Error returns the raw error string
func (err *ClientError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *ClientError) FmtError() string {
	return errorFmt(ClientErrorTypeName, err.Err)
}

/* ----- TRANSIENT ERROR ----- */

// TransientError represents a neo4j-go-driver security error
type TransientError struct {
	Err string
}

// Error returns the raw error string
func (err *TransientError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *TransientError) FmtError() string {
	return errorFmt(TransientErrorTypeName, err.Err)
}

/* ----- SESSION ERROR ----- */

// SessionError represents a neo4j-go-driver security error
type SessionError struct {
	Err string
}

// Error returns the raw error string
func (err *SessionError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *SessionError) FmtError() string {
	return errorFmt(SessionErrorTypeName, err.Err)
}

/* ----- SERVICE UNAVAILABLE ERROR ----- */

// UnavailableError represents a neo4j-go-driver security error
type UnavailableError struct {
	Err string
}

// Error returns the raw error string
func (err *UnavailableError) Error() string {
	return err.Err
}

// FmtError returns the formatted error string
func (err *UnavailableError) FmtError() string {
	return errorFmt(UnavailableErrorTypeName, err.Err)
}
