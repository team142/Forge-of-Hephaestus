package util

func Distance2Dimensions(x1, z1, x2, z2 int) int {
	greaterX, lesserX := OrderIntegers(x1, x2)
	xDiff := greaterX - lesserX

	greaterZ, lesserZ := OrderIntegers(z1, z2)
	zDiff := greaterZ - lesserZ

	result := xDiff + zDiff
	return result

}
