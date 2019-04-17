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

func TestPythonResource(t *testing.T) {
	Convey("resource generator", t, func(c C) {
		targetdir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("resource with request body", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/server_resources/deliveries.raml", apiDef)
			c.So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", targetdir, true, nil, false)

			err = fs.generateResources(targetdir)
			c.So(err, ShouldBeNil)

			// check  api implementation
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "deliveries_api.py"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../fixtures/server_resources/deliveries_api.py")
			c.So(err, ShouldBeNil)
			c.So(s, ShouldEqual, tmpl)
		})

		c.Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
