package currency

import (
	"errors"
	"strconv"
	"strings"
)

// ConvertDolllarsToPennies takes a dollar amount
// as a string, i.e. 1.00, 55.12 etc and converts it
// into an int64
func ConvertDolllarsToPennies(amount string) (int64, error) {
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	// split value on "."
	groups := strings.Split(amount, ".")

	// if there's no . result will still be captured here
	result := groups[0]

	// base string
	r := ""

	// handle the data after the "."
	if len(groups) == 2 {
		if len(groups[1]) != 2 {
			return 0, errors.New("invalid cents")
		}
		r = groups[1]
		if len(r) > 2 {
			r = r[:2]
		}
	}
	// pad with 0, this will be
	// 2 0's if there was no.
	for len(r) < 2 {
		r += "0"
	}

	result += r
	// convert it to an int
	return strconv.ParseInt(result, 10, 64)
}
