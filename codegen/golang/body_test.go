package golang

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/feeltheajf/go-raml/raml"
	"github.com/feeltheajf/go-raml/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructFromBody(t *testing.T) {
	Convey("generate struct body from raml", t, func(c C) {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("simple body", t, func(c C) {
			err = raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			c.So(err, ShouldBeNil)

			s := NewServer(apiDef, "main", "", "examples.com", false, targetDir, nil)
			err := s.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/body"
			typeFiles := []string{
				"UsersIdGetRespBody",
				"UsersPostReqBody",
				"Catanimal",
				"UnionCatanimal",
			}

			for _, f := range typeFiles {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, typeDir, f+".go"))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

			apiFiles := []string{
				"users_api",
				"users_api_IdGet",
				"users_api_IdPut",
				"users_api_Post",
			}

			for _, f := range apiFiles {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, serverAPIDir, "users", f+".go"))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}
		})

		Convey("builtin type doesn't need validation code", t, func(c C) {
			err = raml.ParseFile("../fixtures/struct/validation.raml", apiDef)
			c.So(err, ShouldBeNil)

			s := NewServer(apiDef, "main", "", "examples.com", false, targetDir, nil)
			err := s.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/validation"
			files := []string{
				"builtin_api",
				"builtin_api_Morecomplextype",
				"builtin_api_Scalartype",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, serverAPIDir, "builtin", f+".go"))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}
		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}
