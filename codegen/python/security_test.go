package python

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

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

			err = generateServerSecurity(apiDef.SecuritySchemes, templates(serverKindFlask), targetdir)
			c.So(err, ShouldBeNil)

			// oauth 2 in dropbox
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Dropbox.py"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("./fixtures/security/oauth2_Dropbox.py")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)

			// oauth 2 facebook
			s, err = utils.TestLoadFile(filepath.Join(targetdir, "oauth2_Facebook.py"))
			c.So(err, ShouldBeNil)

			tmpl, err = utils.TestLoadFile("./fixtures/security/oauth2_Facebook.py")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		})

		Convey("routes generation", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/security/dropbox.raml", apiDef)
			c.So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", targetdir, true, nil, false)
			err = fs.generateResources(targetdir)
			c.So(err, ShouldBeNil)

			// check route
			s, err := utils.TestLoadFile(filepath.Join(targetdir, "deliveries_api.py"))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile("./fixtures/security/deliveries_api.py")
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		})

		Convey("flask security classes", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("./fixtures/client/security/client.raml", apiDef)
			c.So(err, ShouldBeNil)

			client := NewClient(apiDef, clientNameRequests, false)
			err = client.generateSecurity(targetdir)
			c.So(err, ShouldBeNil)

			files := []string{
				"oauth2_client_itsyouonline.py",
				"basicauth_client_basic.py",
				"passthrough_client_passthrough.py",
			}
			for _, file := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetdir, file))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join("./fixtures/client/security/flask/", file))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})

		Convey("aiohttp security classes", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("./fixtures/client/security/client.raml", apiDef)
			c.So(err, ShouldBeNil)

			client := NewClient(apiDef, clientNameAiohttp, false)
			err = client.generateSecurity(targetdir)
			c.So(err, ShouldBeNil)

			files := []string{
				"oauth2_client_itsyouonline.py",
				"basicauth_client_basic.py",
				"passthrough_client_passthrough.py",
			}
			for _, file := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetdir, file))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join("./fixtures/client/security/aiohttp/", file))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}

		})
		c.Reset(func() {
			os.RemoveAll(targetdir)
		})
	})
}
