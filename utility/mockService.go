package utility

import "math/rand"

func SimulateAPICall() string {
	result := rand.Intn(100)

	if result < 20 {
		return "FAILED"
	}

	return "SUCCESS"
}
