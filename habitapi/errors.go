package habitapi

type ErrUserNotFound struct {
    Err error
}

func (e *ErrUserNotFound) Error() string {
    return e.Err.Error()
}
