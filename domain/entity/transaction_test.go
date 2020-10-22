package entity

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	transaction, err := NewTransaction(1, 4, 10)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	fmt.Println(transaction.OperationType)
}
