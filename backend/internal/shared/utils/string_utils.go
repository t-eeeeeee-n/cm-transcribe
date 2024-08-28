package utils

import (
	"strings"
	"unicode"
)

// RemoveUnsupportedCharacters AWS Transcribeでサポートされない文字を取り除きます
func RemoveUnsupportedCharacters(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if isSupportedCharacter(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// isSupportedCharacter AWS Transcribeでサポートされている文字かどうかを判定します
func isSupportedCharacter(r rune) bool {
	// ここでは、AWS Transcribeがサポートする範囲の文字のみを許可しています
	// 具体的には、英数字、ハイフン、アポストロフィ、スペース、タブなどを許可します
	return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '-' || r == '\'' || r == ' ' || r == '\t'
}
