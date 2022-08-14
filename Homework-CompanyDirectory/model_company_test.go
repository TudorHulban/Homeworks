package company

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLineage(t *testing.T) {
	company := NewCompany()

	_, errNil := company.FirstManager(nil, nil)
	require.ErrorIs(t, errNil, errNoEmployee)

	ceo := &Employee{
		FullName: "Monica",
	}

	_, errZero := company.FirstManager(nil, nil)
	require.Error(t, errZero)

	require.NoError(t, company.Add(ceo), "adding root employee")

	assert.Equal(t, 1, len(company.team))
	require.Nil(t, company.Hierarchy(ceo))

	m1 := &Employee{
		FullName: "Marcus",
	}

	require.Equal(t, company.Add(m1), errNoManager)

	m1.Manager = ceo

	require.NotNil(t, company.Hierarchy(m1)[0])
	assert.Equal(t, m1, company.Hierarchy(m1)[0])

	w1 := &Employee{
		Manager:  m1,
		FullName: "John Doe",
	}
	assert.NoError(t, company.Add(w1))

	w2 := &Employee{
		Manager:  m1,
		FullName: "Bassem Al Raed",
	}
	assert.NoError(t, company.Add(w2))

	m2 := &Employee{
		Manager:  ceo,
		FullName: "Andre",
	}
	assert.NoError(t, company.Add(m2))

	w3 := &Employee{
		Manager:  m2,
		FullName: "Maurice Ravel",
	}

	assert.NoError(t, company.Add(w3))

	m3 := &Employee{
		Manager:  m1,
		FullName: "Constantin",
	}

	assert.NoError(t, company.Add(m3))

	w4 := &Employee{
		Manager:  m2,
		FullName: "Elena",
	}

	assert.NoError(t, company.Add(w4))

	// testing lineage
	cases := []struct {
		description string
		expectError bool
		a, b        *Employee
		expected    *Employee
	}{
		{"Lineage CEO", false, ceo, ceo, nil},
		{"No lineage but CEO", true, w1, w4, ceo},
		{"Same Team Workers", false, w1, w2, m1},
		{"Board Managers", true, m1, m2, ceo},
	}

	for _, tc := range cases {
		t.Run(tc.description, func(t *testing.T) {
			firstManager, err := company.FirstManager(tc.a, tc.b)

			assert.Equal(t, tc.expectError, err != nil)
			assert.Equal(t, tc.expected, firstManager)
		})
	}
}

func TestList(t *testing.T) {
	// TODO: add testing for List
}
