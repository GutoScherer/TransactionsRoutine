package handler

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/GutoScherer/TransactionsRoutine/domain"
	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
)

type createTransactionHandlerDependencies struct {
	transactionCreator TransactionCreator
}

func TestTransactionCreateHandlerFunc(t *testing.T) {

	logger := log.New(writerMock{}, "", log.LstdFlags)

	type input struct {
		body io.Reader
	}

	dateTime := "2020-10-25T04:37:53.3971639-03:00"
	createdAt, _ := time.Parse(time.RFC3339Nano, dateTime)

	var testCases = []struct {
		testName           string
		input              input
		expectedBodyOutput string
		expectedHTTPCode   int
		dependencies       createTransactionHandlerDependencies
	}{
		{
			testName: "Test empty input body",
			input: input{
				bytes.NewReader([]byte("")),
			},
			expectedBodyOutput: `{"error":"invalid JSON body"}`,
			expectedHTTPCode:   http.StatusBadRequest,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					nil,
				),
			},
		},
		{
			testName: "Test invalid input body",
			input: input{
				bytes.NewReader([]byte("INVALID BODY")),
			},
			expectedBodyOutput: `{"error":"invalid JSON body"}`,
			expectedHTTPCode:   http.StatusBadRequest,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					nil,
				),
			},
		},
		{
			testName: "Test foreign key error",
			input: input{
				bytes.NewReader([]byte(`{"account_id":1, "operation_type_id":1, "amount":100.00}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusUnprocessableEntity,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					repository.ErrForeignKeyConstraint,
				),
			},
		},
		{
			testName: "Test internal error",
			input: input{
				bytes.NewReader([]byte(`{"account_id":1, "operation_type_id":1, "amount":100.00}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusInternalServerError,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					errors.New("unexpected error"),
				),
			},
		},
		{
			testName: "Test invalid data",
			input: input{
				bytes.NewReader([]byte(`{"account_id":1, "operation_type_id":1, "amount":65106510651065106510650}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusUnprocessableEntity,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					repository.ErrInvalidData,
				),
			},
		},

		{
			testName: "Test domain error",
			input: input{
				bytes.NewReader([]byte(`{"account_id":1, "operation_type_id":8, "amount":100}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusUnprocessableEntity,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					nil,
					domain.NewDomainError("Invalid operation type ID"),
				),
			},
		},
		{
			testName: "Test valid input body",
			input: input{
				bytes.NewReader([]byte(`{"account_id":1, "operation_type_id":1, "amount":100.00}`)),
			},
			expectedBodyOutput: fmt.Sprintf(`{"account_id":1,"operation_type":"COMPRA A VISTA","amount":100,"created_at":"%s"}`, `2020-10-25T04:37:53.3971639-03:00`),
			expectedHTTPCode:   http.StatusCreated,
			dependencies: createTransactionHandlerDependencies{
				NewTransactionCreatorMock(
					&presenter.CreateTransactionOutput{
						AccountID:     1,
						OperationType: "COMPRA A VISTA",
						Amount:        100,
						CreatedAt:     createdAt,
					},
					nil,
				),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			rr := httptest.NewRecorder()
			httpHandler := http.HandlerFunc(NewCreateTransaction(tt.dependencies.transactionCreator, logger).HandlerFunc)
			req, err := http.NewRequest("POST", "/transactions", tt.input.body)
			if err != nil {
				t.Error("error to create POST /transactions request")
			}

			httpHandler.ServeHTTP(rr, req)

			var (
				gotHTTPStatusCode = rr.Code
				gotBodyOutput     = strings.TrimRight(rr.Body.String(), "\n")
			)

			if gotHTTPStatusCode != tt.expectedHTTPCode {
				t.Errorf("HTTP Status Code is different from expected '%v', got = '%v'", tt.expectedHTTPCode, gotHTTPStatusCode)
				return
			}

			matchBodyOutput := regexp.MustCompile(tt.expectedBodyOutput).MatchString(gotBodyOutput)
			if !matchBodyOutput {
				t.Errorf("Body Output expected '%v', got = '%v'", tt.expectedBodyOutput, gotBodyOutput)
				return
			}
		})
	}
}
