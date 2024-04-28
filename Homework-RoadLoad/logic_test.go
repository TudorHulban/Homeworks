package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoad(t *testing.T) {
	cases := []struct {
		description    string
		maximumWeight  uint
		vehicleWeights []uint
		expected       uint
	}{
		{"1. empty", 10, []uint{}, 0},
		{"2. one light vehicle", 10, []uint{1}, 0},
		{"3. one light vehicle", 10, []uint{10}, 0},
		{"4. one heavy vehicle", 10, []uint{11}, 1},
		{"5. two light vehicles", 10, []uint{1, 2}, 0},
		{"6. two heavy vehicles", 10, []uint{11, 12}, 2},
		{"7. four vehicles", 10, []uint{11, 12, 1, 1}, 2},
		{"8. four light vehicles", 10, []uint{4, 5, 1, 1}, 0},
		{"9. mixed vehicles", 7, []uint{7, 6, 5, 2, 7, 4, 5, 4}, 5},
		{"10. mixed vehicles", 9, []uint{5, 3, 8, 1, 8, 7, 7, 6}, 4},
	}

	for _, tc := range cases {
		t.Run(tc.description,
			func(t *testing.T) {
				require.EqualValues(t,
					tc.expected,
					Road(tc.maximumWeight, tc.vehicleWeights),
				)
			},
		)
	}
}
