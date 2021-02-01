package util

func BytesToWord(lb, hb uint8) uint16 {
	return uint16(hb)<<8 | uint16(lb)
}

func Btoi(b bool) uint8 {
	if b {
		return 1
	}

	return 0
}
