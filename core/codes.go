package core

const (
	NoError                ErrorCode = 0
	ErrInvalidUri          ErrorCode = 1
	ErrInvalidBody         ErrorCode = 2
	ErrInvalidQuery        ErrorCode = 3
	ErrBindingNotSpecified ErrorCode = 4
	ErrAPIExternalRequest  ErrorCode = 5
	ErrCodeIsNotErrorCode  ErrorCode = 6
)
