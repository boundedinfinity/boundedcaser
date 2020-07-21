package caser

import (
	"fmt"
	"strings"
)

func Convert(input string, options Options) (string, error) {
	var caseType CaseType
	var separator string
	var output string
	var err error

	switch options.Type {
	case "":
		caseType = Camel
	case Camel:
		caseType = Camel
	case LowerSnake:
		caseType = LowerSnake
	case Snake:
		caseType = LowerSnake
	case UpperSnake:
		caseType = UpperSnake
	case Flat:
		caseType = LowerFlat
	case LowerFlat:
		caseType = LowerFlat
	case UpperFlat:
		caseType = UpperFlat
	case Kebab:
		caseType = LowerKebab
	case LowerKebab:
		caseType = LowerKebab
	case UpperKebab:
		caseType = UpperKebab
	case Pascal:
		caseType = Pascal
	case Studly:
		caseType = Pascal
	default:
		return output, fmt.Errorf("")
	}

	switch options.Separator {
	case "":
		separator = " "
	default:
		separator = options.Separator
	}

	switch caseType {
	case Camel:
		output, err = convertToCamel(input, separator)
	case Pascal:
		output, err = convertTo(input, separator, "", strings.Title)
	case LowerSnake:
		output, err = convertTo(input, separator, "_", strings.ToLower)
	case UpperSnake:
		output, err = convertTo(input, separator, "_", strings.ToUpper)
	case LowerFlat:
		output, err = convertTo(input, separator, "", strings.ToLower)
	case UpperFlat:
		output, err = convertTo(input, separator, "", strings.ToUpper)
	case LowerKebab:
		output, err = convertTo(input, separator, "-", strings.ToLower)
	case UpperKebab:
		output, err = convertTo(input, separator, "-", strings.ToUpper)
	}

	return output, err
}

func convertToCamel(input string, sep string) (string, error) {
	words1 := strings.Split(input, sep)

	for i := range words1 {
		if i == 0 {
			words1[i] = strings.ToLower(words1[i])
		} else {
			words1[i] = strings.Title(strings.ToLower(words1[i]))
		}
	}

	return strings.Join(words1, ""), nil
}

func convertTo(input, separator, joiner string, converter func(string) string) (string, error) {
	words := strings.Split(input, separator)

	for i := range words {
		words[i] = converter(words[i])
	}

	return strings.Join(words, joiner), nil
}
