package entity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/domain"
)

func TestNewOperationType(t *testing.T) {
	var testCases = []struct {
		testName       string
		input          uint64
		expectedOutput OperationType
		expectedErr    error
	}{
		{
			testName:       "Test invalid operation type ID",
			input:          5,
			expectedOutput: 0,
			expectedErr:    domain.NewDomainError(fmt.Sprintf("invalid operation type '%d'", 5)),
		},
		{
			testName:       "Test valid operation type ID",
			input:          1,
			expectedOutput: OperationType(1),
			expectedErr:    nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			output, err := NewOperationType(tt.input)
			if err != nil && !reflect.DeepEqual(tt.expectedErr, err) {
				t.Errorf("NewOperationType expected error = '%v', got = '%v'", tt.expectedErr, err)
				return
			}

			if output != 0 && !reflect.DeepEqual(tt.expectedOutput, output) {
				t.Errorf("NewOperationType expected output = '%v', got = '%v'", tt.expectedOutput, output)
				return
			}
		})
	}
}

func TestDebitOperationType(t *testing.T) {
	var testCases = []struct {
		testName       string
		input          uint64
		expectedOutput bool
	}{
		{
			testName:       "Test is not debit operation #1",
			input:          1,
			expectedOutput: true,
		},
		{
			testName:       "Test is not debit operation #2",
			input:          2,
			expectedOutput: true,
		},
		{
			testName:       "Test is not debit operation #3",
			input:          3,
			expectedOutput: true,
		},
		{
			testName:       "Test valid operation type ID",
			input:          4,
			expectedOutput: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			operationType := OperationType(tt.input)

			if operationType.IsDebit() != tt.expectedOutput {
				t.Errorf("OperationType expected output = '%v', got = '%v'", tt.expectedOutput, operationType.IsDebit())
				return
			}
		})
	}
}

func TestOperationTypeDescription(t *testing.T) {
	var testCases = []struct {
		testName       string
		input          uint64
		expectedOutput string
	}{
		{
			testName:       "Test COMPRA A VISTA",
			input:          1,
			expectedOutput: `COMPRA A VISTA`,
		},
		{
			testName:       "Test COMPRA PARCELADA",
			input:          2,
			expectedOutput: `COMPRA PARCELADA`,
		},
		{
			testName:       "Test SAQUE",
			input:          3,
			expectedOutput: `SAQUE`,
		},
		{
			testName:       "Test PAGAMENTO",
			input:          4,
			expectedOutput: `PAGAMENTO`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			operationType := OperationType(tt.input)

			if operationType.String() != tt.expectedOutput {
				t.Errorf("OperationType expected output = '%v', got = '%v'", tt.expectedOutput, operationType.String())
				return
			}
		})
	}
}
