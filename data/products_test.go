package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:   "test_prod",
		Amount: 2,
		SKU:    "abc-dcde-ase",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
