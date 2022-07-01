package maximum_units_on_a_truck

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	var i, maxUnits int
	for truckSize > 0 && i < len(boxTypes) {
		if truckSize >= boxTypes[i][0] {
			maxUnits += boxTypes[i][0] * boxTypes[i][1]
		} else {
			maxUnits += truckSize * boxTypes[i][1]
		}
		truckSize -= boxTypes[i][0]
		i++
	}
	return maxUnits
}
