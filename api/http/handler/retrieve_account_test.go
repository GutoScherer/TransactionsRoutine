package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/GutoScherer/TransactionsRoutine/usecase/presenter"
	"github.com/GutoScherer/TransactionsRoutine/usecase/repository"
	"github.com/gorilla/mux"
)

type retrieveAccountHandlerDependencies struct {
	accountRetriever AccountRetriever
}

func TestAccountRetrieveHandlerFunc(t *testing.T) {

	logger := log.New(writerMock{}, "", log.LstdFlags)

	type input struct {
		accountID string
	}

	var testCases = []struct {
		testName           string
		input              input
		expectedBodyOutput string
		expectedHTTPCode   int
		dependencies       retrieveAccountHandlerDependencies
	}{
		{
			testName: "Test invalid input",
			input: input{
				accountID: "123a",
			},
			expectedBodyOutput: `{"error":"Invalid accountID"}`,
			expectedHTTPCode:   http.StatusBadRequest,
			dependencies: retrieveAccountHandlerDependencies{
				NewAccountRetrieverMock(
					nil,
					nil,
				),
			},
		},
		{
			testName: "Test not found",
			input: input{
				accountID: "123",
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusNotFound,
			dependencies: retrieveAccountHandlerDependencies{
				NewAccountRetrieverMock(
					nil,
					repository.ErrRegisterNotFound,
				),
			},
		},
		{
			testName: "Test internal server error",
			input: input{
				accountID: "100",
			},
			expectedBodyOutput: ``,
			expectedHTTPCode:   http.StatusInternalServerError,
			dependencies: retrieveAccountHandlerDependencies{
				NewAccountRetrieverMock(
					nil,
					errors.New(`unexpected error`),
				),
			},
		},
		{
			testName: "Test valid input",
			input: input{
				accountID: "1",
			},
			expectedBodyOutput: fmt.Sprintf(`{"account_id":1,"document_number":"1231231287"`),
			expectedHTTPCode:   http.StatusOK,
			dependencies: retrieveAccountHandlerDependencies{
				NewAccountRetrieverMock(
					&presenter.RetrieveAccountOutput{
						AccountID:      1,
						DocumentNumber: "1231231287",
					},
					nil,
				),
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%s", tt.input.accountID), nil)
			if err != nil {
				t.Errorf("error to create GET /accounts/%s request", tt.input.accountID)
			}

			router := mux.NewRouter()
			router.HandleFunc("/accounts/{accountID}", NewRetrieveAccount(tt.dependencies.accountRetriever, logger).HandlerFunc).Methods("GET")
			router.ServeHTTP(rr, req)

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
