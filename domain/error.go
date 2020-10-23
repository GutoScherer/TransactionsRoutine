package domain

// Error represents a error of domain
type Error struct {
	message string
}

func (err Error) Error() string {
	return err.message
}

// NewDomainError creates a new domain error struct
func NewDomainError(message string) error {
	return &Error{message: message}
}
