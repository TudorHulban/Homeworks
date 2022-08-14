package exchange

// Exch provides a structure to which a conversion method was added.
type Exch struct{}

const _url = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func (*Exch) ConvertedAmount(amount float64, currencyISO, forDate string) (float64, error) {
	if errValidCurr := validateCurrency(currencyISO); errValidCurr != nil {
		return 0.0, errValidCurr
	}

	if errValidDate := validateDate(forDate); errValidDate != nil {
		return 0.0, errValidDate
	}

	rate, errGet := getRate(_url, currencyISO, forDate)
	if errGet != nil {
		return 0.0, errGet
	}

	return amount * rate, nil
}
