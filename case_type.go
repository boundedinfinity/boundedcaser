package caser

type CaseType string

const (
	// Camel converts "Some Word" to "someWord" (using space separator)
	Camel CaseType = "Camel"

	// Snake converts "Some Word" to "some-word" (using space separator)
	Snake CaseType = "Snake"

	// LowerSnake converts "Some Word" to "some-word" (using space separator)
	//
	// Alias for Snake
	LowerSnake CaseType = "Lower Snake"

	// UpperSnake converts "Some Word" to "SOME-WORD" (using space separator)
	UpperSnake CaseType = "Upper Snake"

	// Flat converts "Some Word" to "someword" (using space separator)
	Flat CaseType = "Flat"

	// LowerFlat converts "Some Word" to "someword" (using space separator)
	//
	// Alias for Flat
	LowerFlat CaseType = "Lower Flat"

	// UpperFlat converts "Some Word" to "SOMEWORD" (using space separator)
	UpperFlat CaseType = "Upper Flat"

	// Kebab converts "Some Word" to "some-word" (using space separator)
	Kebab CaseType = "Kebab"

	// LowerKebab converts "Some Word" to "some-word" (using space separator)
	//
	// Alias for Kebab
	LowerKebab CaseType = "Lower Kebab"

	// UpperKebab converts "Some Word" to "SOME-WORD" (using space separator)
	UpperKebab CaseType = "Upper Kebab"

	// Pascal converts "Some Word" to "SomeWord" (using space separator)
	Pascal CaseType = "Pascal"

	// Studly converts "Some Word" to "SomeWord" (using space separator)
	//
	// Alias for Pascal
	Studly CaseType = "Studly"
)
