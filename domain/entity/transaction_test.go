package entity

import (
	"testing"
)

func TestA(t *testing.T) {
	_, err := NewTransaction(1, 4, 10)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
