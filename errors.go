package dawn

// ErrorType is an unsigned 64-bit error code as defined in the dawn spec.
type ErrorType uint64

// Error represents a error's specification.
type Error struct {
	Err  error
	Type ErrorType
	Meta any
}

type errorMsgs []*Error
