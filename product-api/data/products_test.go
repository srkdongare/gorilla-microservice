package data

import "testing"

func TestValidateFunc(t *testing.T) {
	prod := Product{Name: "sdfg", Description: "-"}

	err := prod.ValidateProduct()

	if err != nil {
		t.Fatal("Validation of Product failed. \n", err)
	}
}
