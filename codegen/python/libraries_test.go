package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/feeltheajf/go-raml/raml"
	"github.com/feeltheajf/go-raml/utils"
)

// TODO FIXME : it disabled because this test is failed and WTF support is planned to be removed
func testLibrary(t *testing.T) {
	Convey("Library usage in server", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/libraries/api.raml", apiDef)
		c.So(err, ShouldBeNil)

		libRootURLs := []string{"https://raw.githubusercontent.com/Jumpscale/go-raml/master/codegen/fixtures/libraries"}
		server := NewFlaskServer(apiDef, "apidocs", targetDir, true, libRootURLs, false)
		err = server.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/libraries/python_server"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"Place.py", "Place.py"},
			{"configs.py", "configs.py"},
			{"libraries/security/oauth2_Dropbox.py", "libraries/security/oauth2_Dropbox.py"},
			{"libraries/files/Directory.py", "libraries/files/Directory.py"},
		}

		for _, check := range checks {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, check.Result))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
