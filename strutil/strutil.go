package strutil

// マルチバイト文字列を含む可能性のある文字列を長さ分切り取る
func Substring(str string, start, length int) string {
	if start < 0 || length <= 0 {
		return str
	}
	r := []rune(str)
	if start+length > len(r) {
		return string(r[start:])
	} else {
		return string(r[start : start+length])
	}
}
