package types

type BadUsageError struct {}

func (bue BadUsageError) Error() string {
	return "Here comes usage..."
}

func (bue *BadUsageError) String() string {
	return bue.Error()
}
