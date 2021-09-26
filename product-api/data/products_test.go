package data

import "testing"

func TestValidateFunc(t *testing.T) {
	prod := Product{Name: "abc", Description: "-"}

	err := prod.ValidateProduct()

	if err != nil {
		t.Fatal("Validation of Product failed.")
	}
}
