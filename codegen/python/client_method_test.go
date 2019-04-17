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

func TestClientMethodWithComplexBody(t *testing.T) {
	Convey("TestClientMethodWithComplexBody", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/body.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameRequests, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/client/complex_body/requests_unmarshall"
		files := []string{
			"arrays_service.py",
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

func TestClientMethodWithQueryParams(t *testing.T) {
	Convey("TestClientMethodWithQueryParams Requests", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/body_with_query_params.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameRequests, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/client/complex_body/query_params_requests/"
		files := []string{
			"animals_service.py",
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

	Convey("TestClientMethodWithQueryParams Aiohttp", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/body_with_query_params.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameAiohttp, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/client/complex_body/query_params_aiohttp/"
		files := []string{
			"animals_service.py",
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

func TestClientMethodWithSpecialChars(t *testing.T) {
	Convey("TestClientMethodWithSpecialChars", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/special_chars.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameRequests, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/special_chars/client"
		files := []string{
			"__init__.py",
			"escape_type_service.py",
			"uri_service.py",
			"User2_0.py",
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

func TestClientMethodWithCatchAllRecursiveURL(t *testing.T) {
	Convey("requests", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameRequests, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/client/requests"
		files := []string{
			"tree_service.py",
			"__init__.py",
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

	Convey("aiohttp", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/catch_all_recursive_url.raml", apiDef)
		c.So(err, ShouldBeNil)

		client := NewClient(apiDef, clientNameAiohttp, true)

		err = client.Generate(targetDir)
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/method/catch_all_recursive_url/client/aiohttp"
		files := []string{
			"tree_service.py",
			"__init__.py",
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
