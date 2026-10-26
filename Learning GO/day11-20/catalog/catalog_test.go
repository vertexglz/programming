package main

import (
	"errors"
	"testing"
)

func TestCatalogAdd(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: 120.00})

	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}

	if len(catalog.Products()) != 1 {
		t.Errorf("len catalog not 1, got:%v", len(catalog.Products()))
	}

}

func TestCatalogAddDuplicate(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: 120.00})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = catalog.Add(Product{Name: "Mouse", Price: 120.00})
	if !errors.Is(err, ErrProductAlreadyExists) {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(catalog.Products()) != 1 {
		t.Errorf("Length is not 1, we got:%v ", len(catalog.Products()))
	}

}

func TestCatalogGet(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: 120.00})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	product, err := catalog.Get("Mouse")

	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if product.Name != "Mouse" {
		t.Errorf("Error, we found Mouse, but got : %v", product.Name)
	}
	if product.Price != 120.00 {
		t.Errorf("Expected price 120, got %v", product.Price)
	}
}

func TestCatalogGetNotFound(t *testing.T) {
	catalog := Catalog{}
	product, err := catalog.Get("Laptop")
	if !errors.Is(err, ErrProductNotFound) {
		t.Fatalf("Unexpected error : %v", err)
	}
	if product != (Product{}) {
		t.Errorf("Expected null , got :%v:%v", product.Name, product.Price)
	}
}

func TestCatalogUpdate(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: 120.00})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	err = catalog.Update("Mouse", Product{Name: "Mouse", Price: 150.00})
	if err != nil {
		t.Fatalf("Update failed : %v", err)
	}
	product, err := catalog.Get("Mouse")

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if product.Name != "Mouse" || product.Price != 150.00 {
		t.Errorf("Expected mouse:150, got : %v : %v", product.Name, product.Price)
	}
}

func TestCatalogUpdateNotFound(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Update("Laptop", Product{Name: "Laptop", Price: 100.00})
	if !errors.Is(err, ErrProductNotFound) {
		t.Fatalf("Unexpected errors: %v", err)
	}

	if len(catalog.Products()) != 0 {
		t.Errorf("expected free catalog, got: %v items", len(catalog.Products()))
	}
}

func TestCatalogRemove(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: 120.00})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	err = catalog.Remove("Mouse")

	if err != nil {
		t.Fatalf("Remove failed: %v", err)
	}

	if len(catalog.Products()) != 0 {
		t.Errorf("Expected 0 items in catalog, got: %d items", len(catalog.Products()))
	}
}

func TestCatalogRemoveNotFound(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Remove("Mouse")
	if !errors.Is(err, ErrProductNotFound) {
		t.Fatalf("Unexpected error : %v", err)
	}
	if len(catalog.Products()) != 0 {
		t.Errorf("Expected 0 items, got : %d items", len(catalog.Products()))
	}
}

func TestCatalogAddInvalidPrice(t *testing.T) {
	catalog := Catalog{}
	err := catalog.Add(Product{Name: "Mouse", Price: -10.00})

	if !errors.Is(err, ErrInvalidProductPrice) {
		t.Fatalf("Unexpected errors:%v", err)
	}
	if len(catalog.Products()) != 0 {
		t.Errorf("expected 0 items, got:%d", len(catalog.Products()))
	}
}
