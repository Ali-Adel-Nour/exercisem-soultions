package affinecipher
import (
	"errors"
	"strings"
	"unicode"
)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func modInverse(a, m int) int {
	a = a % m
	for x := 1; x < m; x++ {
		if (a*x)%m == 1 {
			return x
		}
	}
	return 0
}
func Encode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.New("a and m must be coprime")
	}
	var encoded []rune
	for _, ch := range strings.ToLower(text) {
		if unicode.IsLetter(ch) {
			x := int(ch - 'a')
			y := (a*x + b) % 26
			encoded = append(encoded, rune('a'+y))
		} else if unicode.IsDigit(ch) {
			encoded = append(encoded, ch)
		}
	}
	var groups []string
	for i := 0; i < len(encoded); i += 5 {
		end := i + 5
		if end > len(encoded) {
			end = len(encoded)
		}
		groups = append(groups, string(encoded[i:end]))
	}
	return strings.Join(groups, " "), nil
}
func Decode(text string, a, b int) (string, error) {
	if gcd(a, 26) != 1 {
		return "", errors.New("a and m must be coprime")
	}
	aInv := modInverse(a, 26)
	var result []rune
	for _, ch := range strings.ToLower(text) {
		if unicode.IsLetter(ch) {
			y := int(ch - 'a')
			x := (aInv * (y - b)) % 26
			if x < 0 {
				x += 26
			}
			result = append(result, rune('a'+x))
		} else if unicode.IsDigit(ch) {
			result = append(result, ch)
		}
	}
	return string(result), nil
}
