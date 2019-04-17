package golang

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/codegen/resource"
	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestServerMethodWithSpecialChars(t *testing.T) {
	Convey("TestServerMethodWithSpecialChars", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		c.So(err, ShouldBeNil)

		gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, targetDir, nil)
		_, err = gs.generateServerResources(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/server"
		files := []string{
			filepath.Join(serverAPIDir, "escape_type", "escape_type_api_Post"),
			filepath.Join(serverAPIDir, "uri", "uri_api_Users_idGet"),
			"uri_if",
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

func TestCatchAllRoute(t *testing.T) {
	Convey("TestServerMethodWithSpecialChars", t, func(c C) {
		sm := serverMethod{
			method: &method{},
		}
		sm.Endpoint = "/users/" + resource.CatchAllRoute

		c.So(sm.Route(), ShouldEqual, "/users/"+muxCatchAllRoute)
	})
}

func TestServerMethodWithCatchAllRoute(t *testing.T) {
	Convey("TestServerMethodWithCatchAllRoute", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, targetDir, nil)
		_, err = gs.generateServerResources(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server"
		files := []string{
			"tree_if",
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

func TestServerMethodWithCatchAllRouteInRoot(t *testing.T) {
	Convey("TestServerMethodWithCatchAllRouteInRoot", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_in_root.raml", apiDef)
		c.So(err, ShouldBeNil)

		gs := NewServer(apiDef, "main", "apidocs", "examples.com/libro", true, targetDir, nil)
		err = gs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server-in-root"
		files := []string{
			"path_if",
			"main",
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
