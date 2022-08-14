package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeed(t *testing.T) {
	e := Exch{}

	_, errExch := e.ConvertedAmount(1.00, "RON", "2021-04-13")
	require.Nil(t, errExch, "passed currency is not correct")
}

func TestRateDate(t *testing.T) {
	e := Exch{}

	_, errExch := e.ConvertedAmount(1.00, "RON", "xxx")
	require.Error(t, errExch, "passed date is not correct")

	_, errExch = e.ConvertedAmount(1.00, "RON", "2021-04-13")
	require.Nil(t, errExch)
}

func TestRateCurrency(t *testing.T) {
	e := Exch{}

	_, errExch := e.ConvertedAmount(1.00, "RONI", "2021-04-13")
	require.Error(t, errExch, "passed currency is not correct")
}
