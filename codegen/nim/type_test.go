package nim

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeConversion(t *testing.T) {
	Convey("Test Type Conversion", t, func(c C) {
		Convey("Type conversion", t, func(c C) {
			c.So(toNimType("string", ""), ShouldEqual, "string")
			c.So(toNimType("number", ""), ShouldEqual, "float")
			c.So(toNimType("integer", ""), ShouldEqual, "int")
			c.So(toNimType("boolean", ""), ShouldEqual, "bool")
			c.So(toNimType("file", ""), ShouldEqual, "string")
			c.So(toNimType("date-only", ""), ShouldEqual, "Time")
			c.So(toNimType("time-only", ""), ShouldEqual, "Time")
			c.So(toNimType("datetime", ""), ShouldEqual, "Time")
			c.So(toNimType("Object", ""), ShouldEqual, "Object")
			c.So(toNimType("string[]", ""), ShouldEqual, "seq[string]")
			c.So(toNimType("array", "string"), ShouldEqual, "seq[string]")
			c.So(toNimType("string[][]", ""), ShouldEqual, "seq[seq[string]]")
			//So(toNimType("string | Person"), ShouldEqual, "interface{}")
			//So(toNimType("(string | Person)[]"), ShouldEqual, "[]interface{}")
		})
	})
}
