package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func getValRepetitionCount(strInRunes []rune, idx int) rune {
	valRepetitionCount := '1'
	if idx < len(strInRunes)-1 && unicode.IsDigit(strInRunes[idx+1]) {
		valRepetitionCount = strInRunes[idx+1]
	}

	return valRepetitionCount
}

func isValEscaped(strInRunes []rune, idx int) bool {
	if idx == 0 {
		return false
	}

	isEscaped := false
	for i := 1; i <= idx; i++ {
		if strInRunes[idx-i] == '\\' {
			isEscaped = !isEscaped
		} else {
			break
		}
	}

	return isEscaped
}

func Unpack(str string) (string, error) {
	// early exit
	if str == "" {
		return "", nil
	}

	var strUnpacked strings.Builder

	strInRunes := []rune(str)

	for i, currVal := range str {
		isCurrValDigit := unicode.IsDigit(currVal)
		isCurrValEscaped := isValEscaped(strInRunes, i)

		// check if string starts with number
		if i == 0 && isCurrValDigit {
			return "", ErrInvalidString
		}

		// check if string contains two consecutive numbers
		if i != 0 && isCurrValDigit {
			prevVal := strInRunes[i-1]
			isPrevValDigit := unicode.IsDigit(prevVal)
			isPrevValEscaped := isValEscaped(strInRunes, i-1)

			if isPrevValDigit && !isPrevValEscaped {
				return "", ErrInvalidString
			}
		}

		// check if string contains non-escapable symbol
		if (!isCurrValDigit && currVal != '\\') && isCurrValEscaped {
			return "", ErrInvalidString
		}

		// skip backslash and unescaped number
		if (isCurrValDigit || currVal == '\\') && !isCurrValEscaped {
			continue
		}

		// get current value repetition number
		currValCount := getValRepetitionCount(strInRunes, i)

		for {
			if currValCount == '0' {
				break
			}
			strUnpacked.WriteRune(currVal)
			currValCount--
		}
	}

	return strUnpacked.String(), nil
}
