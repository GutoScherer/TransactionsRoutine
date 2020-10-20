package entity

//Account is an entity for account data storage
type Account struct {
	accountID      uint64
	documentNumber uint64
}

//ID returns the account ID
func (acc Account) ID() uint64 {
	return acc.accountID
}

//DocumentNumber returns the account document number
func (acc Account) DocumentNumber() uint64 {
	return acc.documentNumber
}
