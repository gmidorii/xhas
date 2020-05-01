package xhas

import "strings"

// String returns true if any string in src matches the dst.
func String(dst string, src []string) bool {
	return strAny(dst, src, func(want, target string) bool { return want == target })
}

func strAny(dst string, src []string, operetor func(want, target string) bool) bool {
	for _, s := range src {
		if operetor(s, dst) {
			return true
		}
	}
	return false
}

// StringNotOne returns true if any string in src donot match the dst.
func StringNotOne(dst string, src []string) bool {
	return strNot(dst, src, func(want, target string) bool { return want == target })
}

// StringAll returns true if all string in src match dst.
func StringAll(dst string, src []string) bool {
	return strNot(dst, src, func(want, target string) bool { return want != target })
}

func strNot(dst string, src []string, operetor func(want, target string) bool) bool {
	for _, s := range src {
		if operetor(s, dst) {
			return false
		}
	}
	return true
}

// StringPre returns a string that matches and true if any string in src matches the prefix
func StringPre(prefix string, src []string) (string, bool) {
	return strAnyReturn(prefix, src, strings.HasPrefix)
}

// StringSuf returns a string that matches and true if any string in src matches the suffix
func StringSuf(suffix string, src []string) (string, bool) {
	return strAnyReturn(suffix, src, strings.HasSuffix)
}

// StringPart returns a string that matches and true if any string in src partially matches the dst.
func StringPart(dst string, src []string) (string, bool) {
	return strAnyReturn(dst, src, strings.Contains)
}

func strAnyReturn(dst string, src []string, operetor func(target, want string) bool) (string, bool) {
	for _, s := range src {
		if operetor(s, dst) {
			return s, true
		}
	}
	return "", false
}

// StringPartRune returns a string that matches and true if the rune in src partially matches the dst.
func StringPartRune(dst rune, src []string) (string, bool) {
	for _, s := range src {
		for _, r := range s {
			if dst == r {
				return s, true
			}
		}
	}
	return "", false
}
