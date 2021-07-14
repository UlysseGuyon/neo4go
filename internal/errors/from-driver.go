package errors

const (
	SecurityErrorTypeName    = "Security"
	AuthErrorTypeName        = "Auth"
	ClientErrorTypeName      = "Client"
	TransientErrorTypeName   = "Transient"
	SessionErrorTypeName     = "Session"
	UnavailableErrorTypeName = "Unavailable Service"
)

type SecurityError struct {
	Err string
}

func (err *SecurityError) Error() string {
	return err.Err
}

func (err *SecurityError) FmtError() string {
	return errorFmt(SecurityErrorTypeName, err.Err)
}

type AuthError struct {
	Err string
}

func (err *AuthError) Error() string {
	return err.Err
}

func (err *AuthError) FmtError() string {
	return errorFmt(AuthErrorTypeName, err.Err)
}

type ClientError struct {
	Err string
}

func (err *ClientError) Error() string {
	return err.Err
}

func (err *ClientError) FmtError() string {
	return errorFmt(ClientErrorTypeName, err.Err)
}

type TransientError struct {
	Err string
}

func (err *TransientError) Error() string {
	return err.Err
}

func (err *TransientError) FmtError() string {
	return errorFmt(TransientErrorTypeName, err.Err)
}

type SessionError struct {
	Err string
}

func (err *SessionError) Error() string {
	return err.Err
}

func (err *SessionError) FmtError() string {
	return errorFmt(SessionErrorTypeName, err.Err)
}

type UnavailableError struct {
	Err string
}

func (err *UnavailableError) Error() string {
	return err.Err
}

func (err *UnavailableError) FmtError() string {
	return errorFmt(UnavailableErrorTypeName, err.Err)
}
