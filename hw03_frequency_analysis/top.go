package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	pattern := regexp.MustCompile(`[а-яА-Я\w\d-]+`)
	matched := pattern.FindString(word)
	// `-` is a special case
	if matched == "-" {
		return ""
	}
	return matched
}

func Top10(sentence string) []string {
	wordCounter := make(map[string]int)

	for _, rawWord := range strings.Fields(sentence) {
		word := normalizeWord(rawWord)
		if word == "" {
			continue
		}
		count, ok := wordCounter[word]
		if ok {
			wordCounter[word] = count + 1
		} else {
			wordCounter[word] = 1
		}
	}

	words := make([]string, 0, len(wordCounter))

	for word := range wordCounter {
		words = append(words, word)
	}

	sort.SliceStable(words, func(i, j int) bool {
		leftKey := words[i]
		leftVal := wordCounter[leftKey]
		rightKey := words[j]
		rightVal := wordCounter[rightKey]

		if leftVal == rightVal {
			return leftKey < rightKey
		}

		return leftVal > rightVal
	})

	wordsLength := 10

	if len(wordCounter) < 10 {
		wordsLength = len(wordCounter)
	}

	return words[:wordsLength]
}
