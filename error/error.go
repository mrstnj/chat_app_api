package error

import (
	"fmt"
)

func baseError(errType ErrType, msg string, err error) error {
	return &AppError{
		ErrType: errType,
		Msg:     msg,
		err:     err,
	}
}

func NotFoundError(name string, err error) error {
	return baseError(NotFoundErr, fmt.Sprintf("%s not found", name), err)
}

func DBError(err error) error {
	return baseError(DBErr, "database operation failed", err)
}

func StorageError(err error) error {
	return baseError(StorageErr, "failed to process file storage operation", err)
}

func EmailError(err error) error {
	return baseError(EmailErr, "email delivery failed", err)
}

func APIError(err error) error {
	return baseError(APIErr, "external service request failed", err)
}

func InternalError(err error) error {
	return baseError(InternalErr, "internal server error", err)
}

func InvalidValueError(name string, err error) error {
	return baseError(InvalidValueErr, fmt.Sprintf("invalid %s", name), err)
}

func RequiredError(name string, err error) error {
	return baseError(RequiredValueErr, fmt.Sprintf("required %s", name), err)
}

func AuthError(err error) error {
	return baseError(InvalidValueErr, "authentication failed", err)
}
