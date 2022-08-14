package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAmount(t *testing.T) {
	var exch Exch

	amount, errConvDay1 := exch.ConvertedAmount(1, "RON", "2021-04-13")
	require.NoError(t, errConvDay1, "passed currency is not correct")
	require.Equal(t, 4.9223, amount)

	amount, errConvDay2 := exch.ConvertedAmount(1, "RON", "2021-04-12")
	require.NoError(t, errConvDay2, "passed currency is not correct")
	require.Equal(t, 4.9203, amount)
}
