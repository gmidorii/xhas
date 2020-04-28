package xhas

import "strings"

// String returns true if the string in src matches the dst.
func String(dst string, src []string) bool {
	return hasStr(dst, src, func(want, target string) bool { return want == target })
}

func hasStr(dst string, src []string, operetor func(want, target string) bool) bool {
	for _, s := range src {
		if operetor(dst, s) {
			return true
		}
	}
	return false
}

// StringPre returns a string that matches and true if the string in src matches the prefix
func StringPre(prefix string, src []string) (string, bool) {
	return hasStrReturn(prefix, src, strings.HasPrefix)
}

// StringSuf returns a string that matches and true if the string in src matches the suffix
func StringSuf(suffix string, src []string) (string, bool) {
	return hasStrReturn(suffix, src, strings.HasSuffix)
}

// StringPart returns a string that matches and true if the string in src partially matches the dst.
func StringPart(dst string, src []string) (string, bool) {
	return hasStrReturn(dst, src, strings.Contains)
}

func hasStrReturn(dst string, src []string, operetor func(want, target string) bool) (string, bool) {
	for _, s := range src {
		if operetor(dst, s) {
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
