package caser

import (
	"strings"
)

func Convert(v string, f, t CaseType) string {
	var splitFn func(string) []string
	var mapFn func(string) string
	var joinFn func([]string) string
	var o string

	switch f {
	case CaseType_Camel:
		splitFn = splitOnCapitalOrNumber
	case CaseType_Pascal:
		splitFn = splitOnCapitalOrNumber
	case CaseType_Snake, CaseType_SnakeUpper:
		splitFn = splitOnUnderscore
	case CaseType_Kebab, CaseType_KebabUpper:
		splitFn = splitOnDash
	case CaseType_Phrase:
		splitFn = splitOnSpace
	default:
		o = v
	}

	switch t {
	case CaseType_Camel:
		mapFn = mapPipeline(strings.ToLower, strings.Title)
		joinFn = joinWithNoSpace
	case CaseType_Phrase:
		mapFn = mapPipeline(strings.ToLower, strings.Title)
		joinFn = joinWithSpace
	case CaseType_Snake:
		mapFn = strings.ToLower
		joinFn = joinWithUnderscore
	case CaseType_SnakeUpper:
		mapFn = strings.ToUpper
		joinFn = joinWithUnderscore
	case CaseType_Pascal:
		mapFn = mapPipeline(strings.ToLower, strings.Title)
		joinFn = joinWithNoSpace
	case CaseType_Kebab:
		mapFn = strings.ToLower
		joinFn = joinWithDash
	case CaseType_KebabUpper:
		mapFn = strings.ToUpper
		joinFn = joinWithDash
	default:
		o = v
	}

	if splitFn != nil && mapFn != nil && joinFn != nil {
		o = splitMapJoin(v, splitFn, mapFn, joinFn)

		switch t {
		case CaseType_Camel:
			o = lcFirst(o)
		}
	}

	return o
}

func CamelToPhase(v string) string {
	return Convert(v, CaseType_Camel, CaseType_Phrase)
}

func CamelToSnake(v string) string {
	return Convert(v, CaseType_Camel, CaseType_Snake)
}

func CamelToPascal(v string) string {
	return Convert(v, CaseType_Camel, CaseType_Pascal)
}

func CamelToKebab(v string) string {
	return Convert(v, CaseType_Camel, CaseType_Kebab)
}

func CamelToKebabUpper(v string) string {
	return Convert(v, CaseType_Camel, CaseType_KebabUpper)
}

func PascalToPhase(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_Phrase)
}

func PascalToCamel(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_Camel)
}

func PascalToSnake(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_Snake)
}

func PascalToSnakeUpper(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_SnakeUpper)
}

func PascalToKebab(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_Kebab)
}

func PascalToKebabUpper(v string) string {
	return Convert(v, CaseType_Pascal, CaseType_KebabUpper)
}

func SnakeToPhase(v string) string {
	return Convert(v, CaseType_Snake, CaseType_Phrase)
}

func SnakeToCamel(v string) string {
	return Convert(v, CaseType_Snake, CaseType_Camel)
}

func SnakeToPascal(v string) string {
	return Convert(v, CaseType_Snake, CaseType_Pascal)
}

func SnakeToKebab(v string) string {
	return Convert(v, CaseType_Snake, CaseType_Kebab)
}

func SnakeToKebabUpper(v string) string {
	return Convert(v, CaseType_Snake, CaseType_KebabUpper)
}

func KababToPhase(v string) string {
	return Convert(v, CaseType_Kebab, CaseType_Phrase)
}

func PhraseToCamel(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_Camel)
}

func PhraseToPascal(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_Pascal)
}

func PhraseToSnake(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_Snake)
}

func PhraseToSnakeUpper(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_SnakeUpper)
}

func PhraseToKebab(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_Kebab)
}

func PhraseToKebabUpper(v string) string {
	return Convert(v, CaseType_Phrase, CaseType_KebabUpper)
}
