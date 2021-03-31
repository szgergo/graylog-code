package types

type NoAccessTokenError struct {}

func (nate NoAccessTokenError) Error() string {
	return "No access token was provided"
}

func (nate *NoAccessTokenError) String() string {
	return nate.Error()
}


