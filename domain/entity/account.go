package entity

// Account is an entity for account data storage
type Account struct {
	ID             uint64
	DocumentNumber string
}

// NewAccount creates a new Account struct
func NewAccount(documentNumber string) *Account {
	return &Account{
		DocumentNumber: documentNumber,
	}
}
