package vulkansc

// Error implements the Go error interface for Result.
func (r Result) Error() string {
	return r.String()
}

// Success returns true if the result is >= SUCCESS (meaning not an error).
func (r Result) Success() bool {
	return r >= SUCCESS
}

// ToError returns nil if the result is SUCCESS, otherwise returns the Result as error.
func (r Result) ToError() error {
	if r == SUCCESS {
		return nil
	}
	return r
}
