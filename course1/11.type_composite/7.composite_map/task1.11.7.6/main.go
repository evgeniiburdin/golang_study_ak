package main

import (
	"fmt"
	"strings"
	"unicode"
)

// Word представляет слово и его позицию в предложении
type Word struct {
	Word string
	Pos  int
}

// filterWords Фильтрует текст, заменяя цензурные и повторяющиеся слова
func filterWords(text string, censorMap map[string]string) string {
	// Разделение текста на предложения
	sentences := splitSentences(text)
	if len(sentences) > 1 {
		for i, sentence := range sentences {
			sentences[i] = filterWords(sentence, censorMap)
		}
		return strings.Join(sentences, " ")
	}
	return filterSentence(sentences[0], censorMap)
}

func filterSentence(sentence string, censorMap map[string]string) string {
	// Разделение предложения на слова
	words := strings.Fields(sentence)
	uniqueWordsMap := make(map[string]Word)
	var resultWords []string

	for pos, word := range words {
		lowerWord := strings.ToLower(word)
		if censorReplacement, isCensored := censorMap[lowerWord]; isCensored {
			// Замена цензурного слова
			word = CheckUpper(word, censorReplacement)
		}
		if _, exists := uniqueWordsMap[strings.ToLower(word)]; !exists {
			// Добавление уникального слова в карту
			uniqueWordsMap[strings.ToLower(word)] = Word{Word: strings.ToLower(word), Pos: pos}
			resultWords = append(resultWords, word)
		}
	}

	return WordsToSentence(resultWords)
}

// WordsToSentence Удаляет пустые слова из слайса и объединяет их в предложение
func WordsToSentence(words []string) string {
	filtered := make([]string, 0, len(words))
	for _, word := range words {
		if word != "" {
			filtered = append(filtered, word)
		}
	}
	return strings.ReplaceAll(strings.Join(filtered, " ")+"!", "!!", "!")
}

// CheckUpper Проверяет, нужно ли заменять первую букву на заглавную
func CheckUpper(old, new string) string {
	if len(old) == 0 || len(new) == 0 {
		return new
	}
	chars := []rune(old)
	if unicode.IsUpper(chars[0]) {
		runes := []rune(new)
		new = string([]rune{unicode.ToUpper(runes[0])}) + string(runes[1:])
	}
	return new
}

// splitSentences Разделяет текст на предложения
func splitSentences(message string) []string {
	originSentences := strings.Split(message, "!")
	var sentences []string
	var orphan string

	for i, sentence := range originSentences {
		words := strings.Split(sentence, " ")
		if len(words) == 1 {
			if len(orphan) > 0 {
				orphan += " "
			}
			orphan += words[0] + "!"
			continue
		}
		if orphan != "" {
			originSentences[i] = strings.Join([]string{orphan, sentence}, " ") + "!"
			orphan = ""
		}
		sentences = append(sentences, originSentences[i])
	}

	return sentences
}

func main() {
	text := "Внимание! Внимание! Покупай срочно срочно крипту только у нас! Биткоин лайткоин эфир по низким ценам! " +
		"Беги, беги, успевай стать финансово независимым с помощью крипты! Крипта будущее финансового мира!"
	censorMap := map[string]string{
		"крипта":   "фрукты",
		"крипту":   "фрукты",
		"крипты":   "фруктов",
		"биткоин":  "яблоки",
		"лайткоин": "яблоки",
		"эфир":     "яблоки",
	}
	filteredText := filterWords(text, censorMap)
	fmt.Println(filteredText)
}
