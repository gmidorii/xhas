package xhas

func Int(dst int, src []int) bool {
	for _, s := range src {
		if dst == s {
			return true
		}
	}
	return false
}

func Int16(dst int16, src []int16) bool {
	for _, s := range src {
		if dst == s {
			return true
		}
	}
	return false
}

func Int32(dst int32, src []int32) bool {
	for _, s := range src {
		if dst == s {
			return true
		}
	}
	return false
}

func Int64(dst int64, src []int64) bool {
	for _, s := range src {
		if dst == s {
			return true
		}
	}
	return false
}
