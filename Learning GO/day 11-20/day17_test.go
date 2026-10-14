package main

import "testing"

func TestAddProducts(t *testing.T) {
	products := []Product{
		{Name: "Laptop", Price: 1200.00},
	}

	newProduct := Product{
		Name:  "Laptop",
		Price: 1300,
	}

	got, added := addProducts(products, newProduct)

	if added {
		t.Error("expected added to be false for duplicate product")
	}

	if len(got) != 1 {
		t.Errorf("expected 1 products, got %d", len(got))
	}

	if got[0].Price != 1200.00 {
		t.Errorf("expected product price to be 1200, got %f", got[0].Price)
	}
}
