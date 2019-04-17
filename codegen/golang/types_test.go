package golang

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeConversion(t *testing.T) {
	Convey("Test Type Conversion", t, func(c C) {
		globGoramlPkgDir = "goraml"
		Convey("Type conversion", t, func(c C) {
			c.So(convertToGoType("string", ""), ShouldEqual, "string")
			c.So(convertToGoType("number", ""), ShouldEqual, "float64")
			c.So(convertToGoType("integer", ""), ShouldEqual, "int")
			c.So(convertToGoType("boolean", ""), ShouldEqual, "bool")
			c.So(convertToGoType("file", ""), ShouldEqual, "string")
			c.So(convertToGoType("date-only", ""), ShouldEqual, "goraml.DateOnly")
			c.So(convertToGoType("time-only", ""), ShouldEqual, "goraml.TimeOnly")
			c.So(convertToGoType("Object", ""), ShouldEqual, "Object")
			c.So(convertToGoType("string[]", ""), ShouldEqual, "[]string")
			c.So(convertToGoType("string[][]", ""), ShouldEqual, "[][]string")
			//So(convertToGoType("string | Person"), ShouldEqual, "interface{}")
			//So(convertToGoType("(string | Person)[]"), ShouldEqual, "[]interface{}")
		})
	})
}
