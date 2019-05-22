package errors

func New(msg string) error {
	return &fundamental{
		msg: msg,
	}
}

type fundamental struct {
	msg string
}

func (f *fundamental) Error() string {
	return f.msg
}

func Wrap(err error, msg string) error {
	if err == nil {
		return err
	}

	err = &withCause{
		cause: err,
		msg:   msg,
	}

	return err
}

type withCause struct {
	cause error
	msg   string
}

func (w *withCause) Error() string {
	return w.msg + ": " + w.cause.Error()
}

func (w *withCause) Cause() error {
	return w.cause
}

func Cause(err error) error {
	for {
		causer, ok := err.(causer)
		if !ok {
			break
		}

		err = causer.Cause()
	}

	return err
}

type causer interface {
	Cause() error
}
