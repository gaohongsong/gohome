package utils

// A simple average func to average float list
// and return average float
func Average(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total / float64(len(args))
}
