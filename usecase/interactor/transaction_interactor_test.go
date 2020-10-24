package interactor

import (
	"reflect"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/domain"
	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	usecaseRepository "github.com/GutoScherer/TransactionsRoutine/usecase/repository"
)

type transactionInteractorDependencies struct {
	repo      repository.TransactionRepository
	presenter presenter.TransactionPresenter
}

func TestCreateTransaction(t *testing.T) {
	type input struct {
		accountID       uint64
		operationTypeID uint64
		amount          float64
	}

	var testCases = []struct {
		testName       string
		input          input
		expectedOutput *presenter.CreateTransactionOutput
		expectedErr    error
		dependencies   transactionInteractorDependencies
	}{
		{
			testName: "Test foreign key error",
			input: input{
				accountID:       1,
				operationTypeID: 1,
				amount:          100,
			},
			expectedOutput: nil,
			expectedErr:    usecaseRepository.ErrForeignKeyConstraint,
			dependencies: transactionInteractorDependencies{
				repo:      repository.NewTransactionRepositoryMock(nil, usecaseRepository.ErrForeignKeyConstraint),
				presenter: nil,
			},
		},
		{
			testName: "Test invalid operation type",
			input: input{
				accountID:       1,
				operationTypeID: 6,
				amount:          100,
			},
			expectedOutput: nil,
			expectedErr:    domain.NewDomainError(`new transaction error: invalid operation type '6'`),
			dependencies: transactionInteractorDependencies{
				repo:      repository.NewTransactionRepositoryMock(nil, nil),
				presenter: nil,
			},
		},
		{
			testName: "Test success",
			input: input{
				accountID:       1,
				operationTypeID: 1,
				amount:          100,
			},
			expectedOutput: &presenter.CreateTransactionOutput{
				AccountID:     1,
				OperationType: entity.OperationType(1).String(),
				Amount:        100,
			},
			expectedErr: nil,
			dependencies: transactionInteractorDependencies{
				repo: repository.NewTransactionRepositoryMock(&entity.Transaction{
					AccountID:     1,
					OperationType: 1,
					Amount:        100,
				}, nil),
				presenter: presenter.NewTransactionPresenter(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			interactor := NewTransactionInteractor(tt.dependencies.repo, tt.dependencies.presenter)
			output, err := interactor.Create(tt.input.accountID, tt.input.operationTypeID, tt.input.amount)

			if err != nil && !reflect.DeepEqual(tt.expectedErr, err) {
				t.Errorf("CreateTransaction expected error = '%v', got = '%v'", tt.expectedErr, err)
				return
			}

			if output != nil && !reflect.DeepEqual(tt.expectedOutput, output) {
				t.Errorf("CreateTransaction expected output = '%v', got = '%v'", tt.expectedOutput, output)
				return
			}
		})
	}
}
