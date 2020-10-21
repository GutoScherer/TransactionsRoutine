package entity

// OperationType is an enum that represents the operation types of transactions
//
// Valid values are:
//  1 - Compra à Vista
//  2 - Compra Parcelada
//  3 - Saque
//  4 - Pagamento
type OperationType uint64

const (
	//CompraAVista representa uma compra a vista
	CompraAVista OperationType = 1 + iota
	//CompraParcelada representa uma compra parcelada
	CompraParcelada
	//Saque representa um saque
	Saque
	//Pagamento representa um pagamento
	Pagamento
)

var types = [...]string{
	"COMPRA A VISTA",
	"COMPRA PARCELADA",
	"SAQUE",
	"PAGAMENTO",
}

func (ot OperationType) String() string {
	return types[ot-1]
}

//IsValid checks if the operation type is valid
func (ot OperationType) IsValid() bool {
	return ot <= Pagamento
}

//IsDebit checks if the operation type is a debit operation
func (ot OperationType) IsDebit() bool {
	switch ot {
	case Pagamento:
		return false
	default:
		return true
	}
}
