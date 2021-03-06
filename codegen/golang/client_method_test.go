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

func TestClientMethodWithSpecialChars(t *testing.T) {
	Convey("TestClientMethodWithSpecialChars", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		c.So(err, ShouldBeNil)

		client, err := NewClient(apiDef, "theclient", "examples.com/libro", targetDir, nil)
		c.So(err, ShouldBeNil)

		err = client.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/client"
		files := []string{
			"escape_type_service",
			"uri_service",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f+".go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}

func TestClientMethodCatchAllRecursiveURL(t *testing.T) {
	Convey("TestClientMethodCatchAllRecursiveURL", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		client, err := NewClient(apiDef, "theclient", "examples.com/libro", targetDir, nil)
		c.So(err, ShouldBeNil)

		err = client.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/client"
		files := []string{
			"tree_service",
			"client_the_0_metadata",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f+".go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f+".txt"))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
