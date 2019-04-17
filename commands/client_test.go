package commands

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/utils"
	. "github.com/smartystreets/goconvey/convey"
)

func TestClientGeneration(t *testing.T) {
	Convey("test command client generattion", t, func(c C) {
		targetdir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("Test run client command using go language", t, func(c C) {
			cmd := ClientCommand{
				Language:    "go",
				Dir:         targetdir,
				RamlFile:    "../codegen/fixtures/client_resources/client.raml",
				PackageName: "theclient",
				ImportPath:  "examples.com/client",
			}
			err := cmd.Execute()
			c.So(err, ShouldBeNil)

			s, err := utils.TestLoadFile(filepath.Join(targetdir, "client_structapitest.go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../codegen/golang/fixtures/client_resources/client_structapitest.txt")
			c.So(err, ShouldBeNil)

			c.So(tmpl, ShouldEqual, s)
		})

		c.Reset(func() {
			//cleanup
			os.RemoveAll(targetdir)
		})
	})
}
