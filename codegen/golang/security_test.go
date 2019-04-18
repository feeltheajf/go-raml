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

func TestOauth2Middleware(t *testing.T) {
	Convey("oauth2 middleware", t, func(c C) {

		targetdir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("middleware generation test", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = generateSecurity(apiDef.SecuritySchemes, targetdir, "main")
			c.So(err, ShouldBeNil)

			// oauth 2 facebook
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Facebook_middleware.go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../fixtures/security/oauth2_Facebook_middleware.txt")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)

			// oauth 2 dropbox
			s, err = utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Dropbox_middleware.go"))
			c.So(err, ShouldBeNil)

			tmpl, err = utils.TestLoadFile("../fixtures/security/oauth2_Dropbox_middleware.txt")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)

		})

		Convey("Go routes generation", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox.raml", apiDef)
			c.So(err, ShouldBeNil)

			gs := NewServer(apiDef, "main", "", "examples.com/goraml", true, targetdir, nil)
			_, err = gs.generateServerResources(targetdir)
			c.So(err, ShouldBeNil)

			// check route
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "deliveries_if.go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../fixtures/security/deliveries_if.txt")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		})

		Convey("With included .raml file", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox_with_include.raml", apiDef)
			c.So(err, ShouldBeNil)

			err = generateSecurity(apiDef.SecuritySchemes, targetdir, "main")
			c.So(err, ShouldBeNil)

			// oauth 2 middleware
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "oauth2_DropboxIncluded_middleware.go"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("../fixtures/security/oauth2_DropboxIncluded_middleware.txt")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)

		})

		c.Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
