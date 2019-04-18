package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/feeltheajf/go-raml/raml"
	"github.com/feeltheajf/go-raml/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClass(t *testing.T) {
	Convey("generate python class from raml", t, func(c C) {
		apiDef := new(raml.APIDefinition)
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("python class from raml Types", t, func(c C) {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			c.So(err, ShouldBeNil)

			globAPIDef = apiDef

			_, err = GenerateAllClasses(apiDef, targetDir, false)
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/class/"
			files := []string{
				"EnumCity.py",
				"animal.py",
				"Cage.py",
				"SingleInheritance.py",
				"PlainObject.py",
				"NumberFormat.py",
				"Cat.py",
				"MultipleInheritance.py",
				"EnumString.py",
				"petshop.py",
				"Catanimal.py",
				"UsersIdGetRespBody.py",
				"UsersPostReqBody.py",
				"WithDateTime.py",
				"Tree.py",
				"Alias.py",
				"AliasBuiltin.py",
				"Animal_2_0.py",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
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
