package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestIsEuropeanUnionMember(t *testing.T) {
	tests := map[string]struct {
		countryCode string
		expected    bool
	}{
		"Slovakia": {
			countryCode: "SK",
			expected:    true,
		},
		"Slovenia": {
			countryCode: "SVN",
			expected:    true,
		},
		"Switzerland": {
			countryCode: "CH",
			expected:    false,
		},
		"Italy": {
			countryCode: "ITA",
			expected:    true,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, IsEuropeanUnionMember(test.countryCode))
		})
	}
}
