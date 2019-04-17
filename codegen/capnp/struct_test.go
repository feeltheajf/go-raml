package capnp

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateCapnpSchema(t *testing.T) {
	Convey("generate capnp schema from raml", t, func(c C) {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("./fixtures/struct.raml", &apiDef)
		c.So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("Schema for Python & Nim", t, func(c C) {
			err = GenerateCapnp(&apiDef, targetDir, "nim", "")
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/vanilla"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Animal.capnp", "Animal.capnp"},
				{"Cage.capnp", "Cage.capnp"},
				{"Admin.capnp", "Admin.capnp"},
				{"EnumAdminClearanceLevel.capnp", "EnumAdminClearanceLevel.capnp"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFileRemoveID(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFileRemoveID(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("Schema for Go", t, func(c C) {
			err = GenerateCapnp(&apiDef, targetDir, "go", "main")
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/golang"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"Animal.capnp", "Animal.capnp"},
				{"Cage.capnp", "Cage.capnp"},
				{"Admin.capnp", "Admin.capnp"},
				{"EnumAdminClearanceLevel.capnp", "EnumAdminClearanceLevel.capnp"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFileRemoveID(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFileRemoveID(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
