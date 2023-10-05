package reason

const (
	Success           = "base.success"
	UnknownError      = "base.error.unknown_error"
	DatabaseError     = "base.error.database_error"
	UnauthorizedError = "base.error.unauthorized_error"
	ForbiddenError    = "base.error.forbidden_error"
	RequestBodyError  = "base.error.request_body_error"
)

const (
	UserNotFound         = "error.user.not_found"
	UsernameInvalid      = "error.user.username_invalid"
	UsernameAlreadyExist = "error.user.username_already_exist"
	EmailAlreadyExist    = "error.user.email_already_exist"
	EmailInvalid         = "error.user.email_invalid"
	PasswordInvalid      = "error.user.password_invalid"
	PasswordNotMatch     = "error.user.password_not_match"
)
