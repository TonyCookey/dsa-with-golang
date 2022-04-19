package main

func romanToInt(s string) int {
	numeralsMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
		"L":    50,
		"C":    100,
		"D":    500,
		"M":    1000,
	}

	if value, ok := numeralsMap[s]; ok {
		return value
	}

	var result int
	for i := 0; i < len(s)-1; i++ {
		value := numeralsMap[string(s[i])]
		nextValue := numeralsMap[string(s[i+1])]
		if (value == 1 || value == 10 || value == 100) && value < nextValue {
			result = result - value
			continue
		}
		result += value
	}
	return result + numeralsMap[string(s[len(s)-1])]

}
