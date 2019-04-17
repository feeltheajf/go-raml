package commands

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestServerGeneration(t *testing.T) {
	Convey("test command server generation", t, func(c C) {
		targetdir, err := ioutil.TempDir("", "test_server_command")
		c.So(err, ShouldBeNil)
		Convey("Test run server command using go language", t, func(c C) {

			cmd := ServerCommand{
				Language:    "go",
				Dir:         targetdir,
				RamlFile:    "../codegen/fixtures/server/user_api/api.raml",
				ImportPath:  "examples.com/ramlcode",
				PackageName: "main",
			}
			err := cmd.Execute()
			c.So(err, ShouldBeNil)

			// check users api implementation
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "handlers", "users", "users_api.go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../codegen/fixtures/server/user_api/users_api.txt")
			c.So(err, ShouldBeNil)
			c.So(s, ShouldEqual, tmpl)

			// check user interface
			s, err = utils.TestLoadFile(filepath.Join(targetdir, "users_if.go"))
			c.So(err, ShouldBeNil)

			tmpl, err = utils.TestLoadFile("../codegen/fixtures/server/user_api/users_if.txt")
			c.So(err, ShouldBeNil)
			c.So(s, ShouldEqual, tmpl)

			// check main file
			s, err = utils.TestLoadFile(filepath.Join(targetdir, "main.go"))
			c.So(err, ShouldBeNil)

			tmpl, err = utils.TestLoadFile("../codegen/fixtures/server/user_api/main.txt")
			c.So(err, ShouldBeNil)
			c.So(s, ShouldEqual, tmpl)
		})

		c.Reset(func() {
			//cleanup
			os.RemoveAll(targetdir)
		})
	})
}

func TestServerNoMainGeneration(t *testing.T) {
	Convey("test command server generation without a main", t, func(c C) {
		targetdir, err := ioutil.TempDir("", "test_server_command")
		c.So(err, ShouldBeNil)
		Convey("Test run server command without a main", t, func(c C) {

			cmd := ServerCommand{
				Dir:              targetdir,
				RamlFile:         "../codegen/fixtures/server/user_api/api.raml",
				PackageName:      "main",
				Language:         "go",
				ImportPath:       "examples.com/ramlcode",
				NoMainGeneration: true,
			}
			err := cmd.Execute()
			c.So(err, ShouldBeNil)

			// check main fil
			if _, err := os.Stat(filepath.Join(targetdir, "main.go")); err == nil {
				c.So(errors.New("main.go file exists"), ShouldBeNil)
			}
		})

		c.Reset(func() {
			//cleanup
			os.RemoveAll(targetdir)
		})
	})
}
