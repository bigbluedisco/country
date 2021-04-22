package country

import (
	"testing"
)

func TestByNumeric_Found(t *testing.T) {
	res, ok := ByNumeric(250)
	if !ok || res.Name != "France" {
		t.Fatal(res, ok)
	}
}

func TestByNumeric_NotFound(t *testing.T) {
	res, ok := ByNumeric(0)
	if ok || res.Name != "" {
		t.Fatal(res, ok)
	}
}

func TestByAlpha2_Found(t *testing.T) {
	res, ok := ByAlpha2("FR")
	if !ok || res.Name != "France" {
		t.Fatal(res, ok)
	}
}

func TestByAlpha2_NotFound(t *testing.T) {
	res, ok := ByAlpha2("XX")
	if ok || res.Name != "" {
		t.Fatal(res, ok)
	}
}

func TestByAlpha3_Found(t *testing.T) {
	res, ok := ByAlpha3("FRA")
	if !ok || res.Name != "France" {
		t.Fatal(res, ok)
	}
}

func TestByAlpha3_NotFound(t *testing.T) {
	res, ok := ByAlpha3("XXX")
	if ok || res.Name != "" {
		t.Fatal(res, ok)
	}
}

func TestAll(t *testing.T) {
	all := All()
	if len(all) != 249 {
		t.Fatal("wrong length", len(all))
	}

	if all[0].Name == all[1].Name {
		t.Fatal("duplicates in all slice")
	}
}
