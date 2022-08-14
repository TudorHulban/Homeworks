package exchange

import (
	"time"

	_ "github.com/antchfx/xmlquery"
	"github.com/pkg/errors"
)

func validateCurrency(iso string) error {
	// TODO: add validation that passed is actual ISO symbol.

	if len(iso) != 3 {
		return errors.New("passed currency does has invalid length")
	}

	return nil
}

func validateDate(date string) error {
	layout := "2006-01-02"

	if _, errDate := time.Parse(layout, date); errDate != nil {
		return errors.WithMessagef(errDate, "passed \"%s\" is not a valid date", date)
	}

	// TODO: add validation that passed date in future
	// TODO: add validation that passed date is older than x days

	return nil
}
