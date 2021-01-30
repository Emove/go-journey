/**
 * @author Emove
 * @date 2021/1/27
 */
package util

import "unicode"

func IsCharacter(char rune) bool {
	index := int(char)
	if (index > 32 && index < 48) ||
		(index > 57 && index < 65) ||
		(index > 90 && index < 97) ||
		(index > 122 && index < 127) {
		return true
	}
	return unicode.IsPunct(char)
}
