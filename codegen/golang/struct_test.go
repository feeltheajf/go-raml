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

func TestStruct(t *testing.T) {
	Convey("generate struct from raml", t, func(c C) {
		apiDef := new(raml.APIDefinition)

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("Struct from raml", t, func(c C) {
			err := raml.ParseFile("../fixtures/struct/struct.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = generateStructs(apiDef.Types, targetDir)
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct"
			files := []string{
				"SingleInheritance",
				"MultipleInheritance",
				"ArrayOfCats",
				"BidimensionalArrayOfCats",
				"Petshop",          // using map type & testing case sensitive type name
				"Pet",              // Union
				"ArrayOfPets",      // Array of union
				"Specialization",   // Specialization
				"EnumCity",         // Enum Field
				"Animal",           // using enum
				"EnumString",       // Enum type
				"ValidationString", // validation
				"Dashed",           // field with dash
				"PlainObject",
				"NumberFormat",
				"WithDateTime",
				"Tree",
				"Leaf",
				"Animal_2_0",
				"Dir",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, typeDir, f+".go"))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("With included & inline JSON ", t, func(c C) {
			err := raml.ParseFile("../fixtures/struct/json/api.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = generateStructs(apiDef.Types, targetDir)
			c.So(err, ShouldBeNil)

			err = generateAllStructs(apiDef, targetDir)
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/struct/json"
			files := []string{
				"PersonInclude",
				"PersonPostReqBody",
				"PersonGetRespBody",
				"PersonInType",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, typeDir, f+".go"))
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
