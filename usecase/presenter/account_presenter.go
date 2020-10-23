package presenter

import (
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
)

type accountPresenter struct{}

// AccountPresenter represents the output port for account usecases
type AccountPresenter interface {
	CreateAccountOutput(*entity.Account) *CreateAccountOutput
	RetrieveAccountOutput(*entity.Account) *RetrieveAccountOutput
}

// NewAccountPresenter creates a new AccountPresenter implementation
func NewAccountPresenter() AccountPresenter {
	return &accountPresenter{}
}

func (ap accountPresenter) CreateAccountOutput(acc *entity.Account) *CreateAccountOutput {
	return &CreateAccountOutput{
		AccountID:      acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}
}

func (ap accountPresenter) RetrieveAccountOutput(acc *entity.Account) *RetrieveAccountOutput {
	return &RetrieveAccountOutput{
		AccountID:      acc.ID,
		DocumentNumber: acc.DocumentNumber,
	}
}

// CreateAccountOutput represents the output data for the create account usecase
type CreateAccountOutput struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}

// RetrieveAccountOutput represents the output data for the retrieve account usecase
type RetrieveAccountOutput struct {
	AccountID      uint64 `json:"account_id"`
	DocumentNumber string `json:"document_number"`
}
