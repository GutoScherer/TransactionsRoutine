package entity

import (
	"reflect"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/domain"
)

func TestNewTransaction(t *testing.T) {
	type input struct {
		accountID       uint64
		operationTypeID uint64
		amount          float64
	}

	var testCases = []struct {
		testName       string
		input          input
		expectedOutput *Transaction
		expectedErr    error
	}{
		{
			testName: "Test invalid operation type ID",
			input: input{
				accountID:       1,
				operationTypeID: 5,
				amount:          100,
			},
			expectedOutput: nil,
			expectedErr:    domain.NewDomainError(`new transaction error: invalid operation type '5'`),
		},
		{
			testName: "Test valid debit transaction",
			input: input{
				accountID:       1,
				operationTypeID: 1,
				amount:          100,
			},
			expectedOutput: &Transaction{
				OperationType: 1,
				Account:       Account{ID: 1},
				Amount:        -100,
			},
			expectedErr: nil,
		},
		{
			testName: "Test valid credit transaction",
			input: input{
				accountID:       1,
				operationTypeID: 4,
				amount:          100,
			},
			expectedOutput: &Transaction{
				OperationType: 4,
				Account:       Account{ID: 1},
				Amount:        100,
			},
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			output, err := NewTransaction(tt.input.accountID, tt.input.operationTypeID, tt.input.amount)
			if err != nil && !reflect.DeepEqual(tt.expectedErr, err) {
				t.Errorf("NewTransaction expected error = '%v', got = '%v'", tt.expectedErr, err)
				return
			}

			if output != nil && !reflect.DeepEqual(tt.expectedOutput, output) {
				t.Errorf("NewOperationType expected output = '%v', got = '%v'", tt.expectedOutput, output)
				return
			}
		})
	}
}
