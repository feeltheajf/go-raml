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

func TestMethod(t *testing.T) {
	Convey("server method with display name", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("resource with request body", t, func(c C) {
			apiDef := new(raml.APIDefinition)
			err := raml.ParseFile("../fixtures/server_resources/display_name/api.raml", apiDef)
			c.So(err, ShouldBeNil)

			fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)

			err = fs.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/method/flask/display_name"
			files := []string{
				"coolness_api.py",
			}

			for _, f := range files {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}
		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}

func TestServerMethodWithComplexBody(t *testing.T) {
	Convey("TestServerMethodWithComplexBody", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/body.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/server/complex_body"
		files := []string{
			"arrays_api.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}

func TestServerMethodWithSpecialChars(t *testing.T) {
	Convey("TestServerMethodWithSpecialChars", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/server"
		files := []string{
			"uri_api.py",
			"escape_type_api.py",
			"handlers/escape_type_postHandler.py",
			"handlers/__init__.py",
			"handlers/uri_users_id_getHandler.py",
			"handlers/schema/User2_0_schema.json",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}

func TestServerMethodWithCatchAllRecursiveURL(t *testing.T) {
	Convey("Flask ", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server/flask"
		files := []string{
			"tree_api.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

	Convey("Sanic ", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewSanicServer(apiDef, "apidocs", targetDir, true, nil)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server/sanic"
		files := []string{
			"tree_if.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}

func TestServerMethodWithInRootCatchAllRecursiveURL(t *testing.T) {
	Convey("Flask ", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_in_root.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server/flask-in-root"
		files := []string{
			"path_api.py",
			"app.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

	Convey("Sanic ", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_in_root.raml", apiDef)
		c.So(err, ShouldBeNil)

		fs := NewSanicServer(apiDef, "apidocs", targetDir, true, nil)

		err = fs.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/server/sanic-in-root"
		files := []string{
			"path_if.py",
			"app.py",
		}

		for _, f := range files {
			s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
			c.So(err, ShouldBeNil)

			tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
			c.So(err, ShouldBeNil)

			c.So(s, ShouldEqual, tmpl)
		}

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}
