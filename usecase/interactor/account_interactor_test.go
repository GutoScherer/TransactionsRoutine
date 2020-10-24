package interactor

import (
	"reflect"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/domain/entity"
	"github.com/GutoScherer/TransactionsRoutine/domain/repository"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	usecaseRepository "github.com/GutoScherer/TransactionsRoutine/usecase/repository"
)

type accountInteractorDependencies struct {
	repo      repository.AccountRepository
	presenter presenter.AccountPresenter
}

func TestCreateAccount(t *testing.T) {
	type input struct {
		documentNumber string
	}

	var testCases = []struct {
		testName       string
		input          input
		expectedOutput *presenter.CreateAccountOutput
		expectedErr    error
		dependencies   accountInteractorDependencies
	}{
		{
			testName:       "Test duplicated entry",
			input:          input{documentNumber: "12312312387"},
			expectedOutput: nil,
			expectedErr:    usecaseRepository.ErrDuplicatedEntry,
			dependencies: accountInteractorDependencies{
				repo:      repository.NewAccountRepositoryMock(nil, usecaseRepository.ErrDuplicatedEntry),
				presenter: presenter.NewAccountPresenter(),
			},
		},
		{
			testName:       "Test invalid data",
			input:          input{documentNumber: "12312312387684068406540584068406845065406840648068406840"},
			expectedOutput: nil,
			expectedErr:    usecaseRepository.ErrInvalidData,
			dependencies: accountInteractorDependencies{
				repo:      repository.NewAccountRepositoryMock(nil, usecaseRepository.ErrInvalidData),
				presenter: presenter.NewAccountPresenter(),
			},
		},
		{
			testName:       "Test success",
			input:          input{documentNumber: "00000000191"},
			expectedOutput: &presenter.CreateAccountOutput{DocumentNumber: "00000000191"},
			expectedErr:    nil,
			dependencies: accountInteractorDependencies{
				repo:      repository.NewAccountRepositoryMock(entity.NewAccount("00000000191"), nil),
				presenter: presenter.NewAccountPresenter(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			interactor := NewAccountInteractor(tt.dependencies.repo, tt.dependencies.presenter)
			output, err := interactor.Create(tt.input.documentNumber)

			if err != nil && !reflect.DeepEqual(tt.expectedErr, err) {
				t.Errorf("CreateAccount expected error = '%v', got = '%v'", tt.expectedErr, err)
				return
			}

			if output != nil && !reflect.DeepEqual(tt.expectedOutput, output) {
				t.Errorf("CreateAccount expected output = '%v', got = '%v'", tt.expectedOutput, output)
				return
			}
		})
	}
}

func TestRetrieveAccount(t *testing.T) {
	type input struct {
		accountID uint64
	}

	var testCases = []struct {
		testName       string
		input          input
		expectedOutput *presenter.RetrieveAccountOutput
		expectedErr    error
		dependencies   accountInteractorDependencies
	}{
		{
			testName:       "Test not found",
			input:          input{accountID: 1},
			expectedOutput: nil,
			expectedErr:    usecaseRepository.ErrRegisterNotFound,
			dependencies: accountInteractorDependencies{
				repo:      repository.NewAccountRepositoryMock(nil, usecaseRepository.ErrRegisterNotFound),
				presenter: presenter.NewAccountPresenter(),
			},
		},
		{
			testName:       "Test success",
			input:          input{accountID: 1},
			expectedOutput: &presenter.RetrieveAccountOutput{DocumentNumber: "00000000191"},
			expectedErr:    nil,
			dependencies: accountInteractorDependencies{
				repo:      repository.NewAccountRepositoryMock(entity.NewAccount("00000000191"), nil),
				presenter: presenter.NewAccountPresenter(),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			interactor := NewAccountInteractor(tt.dependencies.repo, tt.dependencies.presenter)
			output, err := interactor.RetrieveByID(tt.input.accountID)

			if err != nil && !reflect.DeepEqual(tt.expectedErr, err) {
				t.Errorf("CreateAccount expected error = '%v', got = '%v'", tt.expectedErr, err)
				return
			}

			if output != nil && !reflect.DeepEqual(tt.expectedOutput, output) {
				t.Errorf("CreateAccount expected output = '%v', got = '%v'", tt.expectedOutput, output)
				return
			}
		})
	}
}
