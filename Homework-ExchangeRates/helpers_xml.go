package exchange

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/antchfx/xmlquery"
)

func fetchData(url string) (*xmlquery.Node, error) {
	return xmlquery.LoadURL(url)
}

func findDayData(data *xmlquery.Node, day string) *xmlquery.Node {
	daysData := xmlquery.Find(data, "//Cube")

	for _, v := range daysData {
		if v.SelectAttr("time") == day {
			//fmt.Println(v.SelectAttr("time"))

			return v
		}
	}

	return nil
}

func findDayRate(dayData *xmlquery.Node, currencyISO string) (float64, error) {
	dayRates := xmlquery.Find(dayData, "//Cube")

	if len(dayRates) == 0 {
		return 0.0, errors.New("no data found")
	}

	for _, v := range dayRates {
		if v.SelectAttr("currency") == currencyISO {
			//fmt.Println(v.SelectAttr("rate"))

			res, errConv := strconv.ParseFloat(v.SelectAttr("rate"), 10)
			if errConv != nil {
				return 0.00, nil
			}

			return res, nil
		}
	}

	return 0.00, nil
}

func getRate(url, currencyISO, forDate string) (float64, error) {
	data, err := fetchData(url)
	if err != nil {
		return 0.0, err
	}

	if data == nil {
		return 0.0, fmt.Errorf("no data found for date: %s", forDate)
	}

	day := findDayData(data, forDate)
	if day == nil {
		return 0.0, fmt.Errorf("no day data found for date: %s", forDate)
	}

	return findDayRate(day, currencyISO)
}
