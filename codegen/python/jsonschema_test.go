package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"
)

func TestJSONSchema(t *testing.T) {
	Convey("JSON Schema test", t, func(c C) {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("array", t, func(c C) {
			err := raml.ParseFile("../fixtures/raml-examples/typesystem/array-type.raml", apiDef)
			c.So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", targetDir, true, nil)
			err = server.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/array_type"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Email_schema.json", "Email_schema.json"},
				{"EmailsLong_schema.json", "EmailsLong_schema.json"},
				{"EmailsShort_schema.json", "EmailsShort_schema.json"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, handlersDir, jsonSchemaDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("simple", t, func(c C) {
			err := raml.ParseFile("../fixtures/raml-examples/typesystem/simple.raml", apiDef)
			c.So(err, ShouldBeNil)

			server := NewSanicServer(apiDef, "apidocs", targetDir, true, nil)
			err = server.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/simple"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"User_schema.json", "User_schema.json"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, handlersDir, jsonSchemaDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("struct", t, func(c C) {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = generateJSONSchema(apiDef, targetDir)
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/json_schema/struct"
			files := []string{
				"EnumCity_schema.json",
				"animal_schema.json",
				"Cage_schema.json",
				"SingleInheritance_schema.json",
				"Dashed_schema.json",
				"PlainObject_schema.json",
				"NumberFormat_schema.json",
				"Cat_schema.json",
				"MultipleInheritance_schema.json",
				//"EnumString_schema.json",
				"Catanimal_schema.json",
				"UsersPostReqBody_schema.json",
				"Alias_schema.json",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, jsonSchemaDir, f))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
