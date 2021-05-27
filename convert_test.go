package caser_test

import (
	"testing"

	"github.com/boundedinfinity/caser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLoader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Caser Suite")
}

var (
	aPhrase    = "Some Test 4 Thing"
	camel      = "someTest4Thing"
	pascal     = "SomeTest4Thing"
	snakeLower = "some_test_4_thing"
	snakeUpper = "SOME_TEST_4_THING"
	kababLower = "some-test-4-thing"
	kababUpper = "SOME-TEST-4-THING"
)

var _ = Describe("Smoke Test", func() {
	It("Phrase to Camel", func() {
		actual := caser.PhraseToCamel(aPhrase)
		Expect(actual).To(Equal(camel))
	})

	It("Phrase to Pascal", func() {
		actual := caser.PhraseToPascal(aPhrase)
		Expect(actual).To(Equal(pascal))
	})

	It("Phrase to Snake (lower)", func() {
		actual := caser.PhraseToSnake(aPhrase)
		Expect(actual).To(Equal(snakeLower))
	})

	It("Phrase to Snake (upper)", func() {
		actual := caser.PhraseToSnakeUpper(aPhrase)
		Expect(actual).To(Equal(snakeUpper))
	})

	It("Phrase to Kabab (lower)", func() {
		actual := caser.PhraseToKebab(aPhrase)
		Expect(actual).To(Equal(kababLower))
	})

	It("Phrase to Kabab (upper)", func() {
		actual := caser.PhraseToKebabUpper(aPhrase)
		Expect(actual).To(Equal(kababUpper))
	})

	It("Camel to Phrase", func() {
		actual := caser.CamelToPhase(camel)
		Expect(actual).To(Equal(aPhrase))
	})

	It("Pascal to Phrase", func() {
		actual := caser.CamelToPhase(pascal)
		Expect(actual).To(Equal(aPhrase))
	})

	It("Snake to Phrase", func() {
		actual := caser.SnakeToPhase(snakeLower)
		Expect(actual).To(Equal(aPhrase))

		actual = caser.SnakeToPhase(snakeUpper)
		Expect(actual).To(Equal(aPhrase))
	})

	It("Kabab to Phrase", func() {
		actual := caser.KababToPhase(kababLower)
		Expect(actual).To(Equal(aPhrase))

		actual = caser.KababToPhase(kababUpper)
		Expect(actual).To(Equal(aPhrase))
	})
})
