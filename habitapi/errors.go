package habitapi

type ErrResourceNotFound struct {
    Err error
}

func (e *ErrResourceNotFound) Error() string {
    return e.Err.Error()
}
