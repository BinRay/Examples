package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Given some integer with a starting value", t, func() {
		//the statement x := 1 will be executed twice.
		//This is because it appears once in each of the top-level "Convey" function calls.
		//Each "Convey" function call creates a new test context,
		//within which x := 1 is executed.
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})

		Convey("When the integer is decremented", func() {
			x--

			Convey("The value should be less by one", func() {
				So(x, ShouldEqual, 0)
			})
		})
	})
}
