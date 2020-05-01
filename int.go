package xhas

func Int(dst int, src []int) bool {
	return intAny(dst, src, func(want, target int) bool { return target == want })
}

func IntOdd(src []int) bool {
	return intAny(2, src, func(want, target int) bool { return (target % want) == 1 })
}

func IntEven(src []int) bool {
	return intAny(2, src, func(want, target int) bool { return (target % want) == 0 })
}

func intAny(dst int, src []int, ope func(want, target int) bool) bool {
	for _, s := range src {
		if ope(dst, s) {
			return true
		}
	}
	return false
}

func IntOddAll(src []int) bool {
	return intNot(2, src, func(want, target int) bool { return (target % want) != 1 })
}

func IntEvenAll(src []int) bool {
	return intNot(2, src, func(want, target int) bool { return (target % want) != 0 })
}

func intNot(dst int, src []int, ope func(want, target int) bool) bool {
	for _, s := range src {
		if ope(dst, s) {
			return false
		}
	}
	return true
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
