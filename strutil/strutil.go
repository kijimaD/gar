package strutil

import (
	"github.com/manifoldco/promptui"
)

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

// yes/noを返す
func YorN(b bool) string {
	if b {
		return "Yes"
	} else {
		return "No"
	}
}

// プロンプトを取得する
func GetPrompt() *promptui.Select {
	prompt := promptui.Select{
		Label: "Send reply[yes/no]",
		Items: []string{"yes", "no"},
	}
	return &prompt
}
