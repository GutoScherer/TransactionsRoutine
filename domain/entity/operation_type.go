package entity

//OperationType represents the operation type of a transaction
type OperationType struct {
	operationTypeID uint64
	description     string
}

//ID returns the operation type ID
func (ot OperationType) ID() uint64 {
	return ot.operationTypeID
}

//Description returns the operation type description
func (ot OperationType) Description() string {
	return ot.description
}
