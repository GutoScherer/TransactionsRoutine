package entity

import "time"

//Transaction is an entity for transactions data storage
type Transaction struct {
	transactionID uint64
	account       *Account
	operationType *OperationType
	amount        float64
	eventDate     time.Time
}

//ID returns the transaction ID
func (t Transaction) ID() uint64 {
	return t.transactionID
}

//Account returns the account of the transaction
func (t Transaction) Account() *Account {
	return t.account
}

//OperationType returns the transaction's operation type
func (t Transaction) OperationType() *OperationType {
	return t.operationType
}

//Amount returns the amount of the transaction
func (t Transaction) Amount() float64 {
	return t.amount
}

//EventDate returns the time that the transaction ocurred
func (t Transaction) EventDate() time.Time {
	return t.eventDate
}
