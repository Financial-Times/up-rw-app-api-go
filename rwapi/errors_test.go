package rwapi

import (
	"testing"
)

func TestConstraintOrTransactionSimpleError1(t *testing.T) {
	e := ConstraintOrTransactionError{"foo", nil}
	expected := `foo`
	if e.Error() != expected {
		t.Errorf("expected %s but got %s", expected, e.Error())
	}
}

func TestConstraintOrTransactionSimpleError2(t *testing.T) {
	e := ConstraintOrTransactionError{"foo", []string{}}
	expected := `foo`
	if e.Error() != expected {
		t.Errorf("expected %s but got %s", expected, e.Error())
	}
}

func TestConstraintOrTransactionDetailedError(t *testing.T) {
	e := ConstraintOrTransactionError{"foo", []string{"bar", "baz"}}
	expected := `foo
details: {
	'bar'
	'baz'
}
`
	if e.Error() != expected {
		t.Errorf("expected %s but got %s", expected, e.Error())
	}
}
