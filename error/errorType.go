package error

type ErrType int

const (
	NotFoundErr ErrType = iota
	DBErr
	StorageErr
	EmailErr
	APIErr
	InternalErr
	InvalidValueErr
	RequiredValueErr
)

type AppError struct {
	ErrType ErrType
	Msg     string
	err     error
}

func (e AppError) Error() string {
	return e.err.Error()
}

func (e AppError) Unwrap() error {
	return e.err
}
