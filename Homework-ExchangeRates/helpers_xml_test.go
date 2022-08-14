package exchange

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ISO = "RON"

func TestGetRate(t *testing.T) {
	rate, errGet := getRate(_url, _ISO, "2022-08-13")
	require.NoError(t, errGet)

	fmt.Println("rate:", rate)

}
