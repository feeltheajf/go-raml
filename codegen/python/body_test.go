package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClassFromBody(t *testing.T) {
	Convey("Class from method body", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("from RAML", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			c.So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)
			err = fs.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/from_body/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"handlers/schema/UsersPostReqBody_schema.json", "UsersPostReqBody_schema.json"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("from RAML with JSON", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			c.So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)
			err = fs.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/from_body/json/"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"handlers/schema/PersonPostReqBody_schema.json", "PersonPostReqBody_schema.json"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
