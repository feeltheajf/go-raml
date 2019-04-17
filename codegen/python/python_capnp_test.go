package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"fmt"

	"github.com/feeltheajf/go-raml/raml"
	"github.com/feeltheajf/go-raml/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGeneratePythonCapnpClasses(t *testing.T) {
	Convey("generate python class from raml", t, func(c C) {
		apiDef := new(raml.APIDefinition)
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("python class from raml Types", t, func(c C) {
			err := raml.ParseFile("./fixtures/python_capnp/types.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = GeneratePythonCapnpClasses(apiDef, targetDir)
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/python_capnp/"
			files := []string{
				"EnumCity.%s",
				"Animal.%s",
				"Cage.%s",
				"SingleInheritance.%s",
				"PlainObject.%s",
				"NumberFormat.%s",
				"Cat.%s",
				"MultipleInheritance.%s",
				"Petshop.%s",
				"WithDateTime.%s",
				"EnumEnumCityEnumParks.%s",
			}

			for _, f := range files {
				// check the python classes
				class := fmt.Sprintf(f, "py")
				s, err := utils.TestLoadFile(filepath.Join(targetDir, class))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, class))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)

				// check the capnp schemas
				schema := fmt.Sprintf(f, "capnp")
				s, err = utils.TestLoadFileRemoveID(filepath.Join(targetDir, schema))
				c.So(err, ShouldBeNil)

				tmpl, err = utils.TestLoadFileRemoveID(filepath.Join(rootFixture, schema))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

			// make sure these files are not exists
			filesNotExist := []string{
				"Alias.capnp",     // alias of non builtin type, use .capnp of aliased type
				"AliasBuiltin.py", // no support for builtin type
			}
			for _, f := range filesNotExist {
				_, err := os.Stat(filepath.Join(targetDir, f))
				c.So(err, ShouldNotBeNil)
			}

		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})

	})
}
