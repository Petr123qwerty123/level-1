package tasks

func Task8(x int64, i uint8, bitValue uint8) int64 {
	mask := int64(1) << i

	if bitValue == 1 {
		x |= mask
	} else {
		x &= ^mask
	}

	return x
}
