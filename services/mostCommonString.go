package services

func MostCommonString(arr []string) (string, int) {
	if len(arr) == 0 {
		return "", 0
	}

	freqs := make(map[string]int, len(arr)/10)

	maxCount := 0
	mostCommon := ""

	for _, val := range arr {
		if val == "" {
			continue
		}

		freqs[val]++

		if freqs[val] > maxCount {
			maxCount = freqs[val]
			mostCommon = val
		}
	}
	return mostCommon, maxCount
}
