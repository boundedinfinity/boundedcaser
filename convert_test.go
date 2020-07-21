package caser_test

import (
	"testing"

	"github.com/boundedinfinity/caser"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec(t *testing.T) {
	input1 := "A test String"
	expectedCamelLower := "aTestString"
	expectedSnakeLower := "a_test_string"
	expectedSnakeUpper := "A_TEST_STRING"
	expectedKebabLower := "a-test-string"
	expectedKebabUpper := "A-TEST-STRING"
	expectedFlatLower := "ateststring"
	expectedFlatUpper := "ATESTSTRING"
	expectedPascal := "ATestString"

	Convey("Caser", t, func() {
		Convey("Convert()", func() {
			Convey("empty (defaults to Camel)", func() {
				actual, err := caser.Convert(input1, caser.Options{})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedCamelLower)
			})

			Convey("Camel", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Camel,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedCamelLower)
			})

			Convey("Snake", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Snake,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedSnakeLower)
			})

			Convey("LowerSnake", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.LowerSnake,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedSnakeLower)
			})

			Convey("UpperSnake", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.UpperSnake,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedSnakeUpper)
			})

			Convey("Kebab", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Kebab,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedKebabLower)
			})

			Convey("LowerKebab", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.LowerKebab,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedKebabLower)
			})

			Convey("UpperKebab", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.UpperKebab,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedKebabUpper)
			})

			Convey("Flat", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Flat,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedFlatLower)
			})

			Convey("LowerFlat", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.LowerFlat,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedFlatLower)
			})

			Convey("UpperFlat", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.UpperFlat,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedFlatUpper)
			})

			Convey("Pascal", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Pascal,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedPascal)
			})

			Convey("Studly", func() {
				actual, err := caser.Convert(input1, caser.Options{
					Type: caser.Studly,
				})

				So(err, ShouldBeNil)
				So(actual, ShouldEqual, expectedPascal)
			})
		})
	})
}
