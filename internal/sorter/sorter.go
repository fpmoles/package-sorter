package sorter

import "errors"

const (
	stackStandard = "STANDARD"
	stackSpecial  = "SPECIAL"
	stackRejected = "REJECTED"
)

// Sort dispatches a package to the correct stack based on its dimensions and mass.
// Dimensions are in centimeters; mass is in kilograms.
// Returns an error if any input value is not positive.
func Sort(width, height, length, mass float64) (string, error) {
	if width <= 0 || height <= 0 || length <= 0 || mass <= 0 {
		return "", errors.New("width, height, length, and mass must all be positive")
	}

	bulky := width*height*length >= 1_000_000 || width >= 150 || height >= 150 || length >= 150
	heavy := mass >= 20

	switch {
	case bulky && heavy:
		return stackRejected, nil
	case bulky || heavy:
		return stackSpecial, nil
	default:
		return stackStandard, nil
	}
}
