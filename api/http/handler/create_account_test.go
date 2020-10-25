package handler

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
)

type createAccountHandlerDependencies struct {
	accountCreator AccountCreator
}

func TestAccountCreateHandlerFunc(t *testing.T) {

	logger := log.New(writerMock{}, "", log.LstdFlags)

	type input struct {
		body io.Reader
	}

	var testCases = []struct {
		testName           string
		input              input
		expectedBodyOutput string
		expectedHTTPCode   int
		dependencies       createAccountHandlerDependencies
	}{
		{
			testName: "Test empty input body",
			input: input{
				bytes.NewReader([]byte("")),
			},
			expectedBodyOutput: `{"error":"invalid JSON body"}`,
			expectedHTTPCode:   http.StatusBadRequest,
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
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
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
					nil,
					nil,
				),
			},
		},
		{
			testName: "Test duplicated entry",
			input: input{
				bytes.NewReader([]byte(`{"document_number":"12312312387"}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusConflict,
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
					nil,
					repository.ErrDuplicatedEntry,
				),
			},
		},
		{
			testName: "Test internal error",
			input: input{
				bytes.NewReader([]byte(`{"document_number":"123123123871231231238712312312387123123123871231231238712312312387"}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusInternalServerError,
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
					nil,
					errors.New("unexpected error"),
				),
			},
		},
		{
			testName: "Test invalid data",
			input: input{
				bytes.NewReader([]byte(`{"document_number":"123123123871231231238712312312387123123123871231231238712312312387"}`)),
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusUnprocessableEntity,
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
					nil,
					repository.ErrInvalidData,
				),
			},
		},
		{
			testName: "Test valid input body",
			input: input{
				bytes.NewReader([]byte(`{"document_number":"12312312387"}`)),
			},
			expectedBodyOutput: `{"account_id":1,"document_number":"12312312387"}`,
			expectedHTTPCode:   http.StatusCreated,
			dependencies: createAccountHandlerDependencies{
				NewAccountCreatorMock(
					&presenter.CreateAccountOutput{
						AccountID:      1,
						DocumentNumber: "12312312387",
					},
					nil,
				),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			rr := httptest.NewRecorder()
			httpHandler := http.HandlerFunc(NewCreateAccount(tt.dependencies.accountCreator, logger).HandlerFunc)
			req, err := http.NewRequest("POST", "/accounts", tt.input.body)
			if err != nil {
				t.Error("error to create POST /accounts request")
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
				t.Errorf("Body Output expected '%v', got = %v", tt.expectedBodyOutput, gotBodyOutput)
				return
			}
		})
	}
}
