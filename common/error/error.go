package error

type CommonErr struct {
	SomeString string
}

func (C *CommonErr) Error() string {
	return C.SomeString + " is Error."
}
