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

var (
	testLibRootURLs = []string{
		"https://raw.githubusercontent.com/Jumpscale/go-raml/master/codegen/fixtures/libraries",
		"https://raw.githubusercontent.com/Jumpscale/go-raml/libraries-in-file/codegen/fixtures/libraries/",
	}
)

func TestLibrary(t *testing.T) {
	Convey("Library usage in server", t, func(c C) {
		var apiDef raml.APIDefinition

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		err = raml.ParseFile("../fixtures/libraries/api.raml", &apiDef)
		c.So(err, ShouldBeNil)

		server := NewServer(&apiDef, "main", "apidocs", "examples.com/ramlcode", true,
			targetDir, testLibRootURLs)
		err = server.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "../fixtures/libraries/go_server"
		checks := []struct {
			Result   string
			Expected string
		}{
			{filepath.Join(typeDir, "Place.go"), "Place.txt"},
			{filepath.Join(serverAPIDir, "dirs", "dirs_api.go"), "dirs_api.txt"},
			{filepath.Join(serverAPIDir, "configs", "configs_api.go"), "configs_api.txt"},
			{filepath.Join(serverAPIDir, "configs", "configs_api_Put.go"), "configs_api_Put.txt"},
			{filepath.Join(serverAPIDir, "configs", "configs_api_Post.go"), "configs_api_Post.txt"},
			{filepath.Join(serverAPIDir, "configs", "configs_api_Get.go"), "configs_api_Get.txt"},
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

	Convey("Library usage in client", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/libraries/api.raml", apiDef)
		c.So(err, ShouldBeNil)

		client, err := NewClient(apiDef, "theclient", "examples.com/theclient", targetDir, testLibRootURLs)
		c.So(err, ShouldBeNil)

		err = client.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "../fixtures/libraries/go_client"
		checks := []struct {
			Result   string
			Expected string
		}{
			{filepath.Join(typeDir, "Place.go"), "Place.txt"},
			{"client_exampleapi.go", "client_exampleapi.txt"},
			{"client_utils.go", "client_utils.txt"},
			{"dirs_service.go", "dirs_service.txt"},
			{"configs_service.go", "configs_service.txt"},
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

	Convey("raml-examples", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("server", t, func(c C) {

			var apiDef raml.APIDefinition

			err = raml.ParseFile("../fixtures/raml-examples/libraries/api.raml", &apiDef)
			c.So(err, ShouldBeNil)

			server := NewServer(&apiDef, "main", "apidocs", "examples.com/libro", true, targetDir, nil)
			err = server.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "../fixtures/libraries/raml-examples/go_server"
			checks := []struct {
				Result   string
				Expected string
			}{
				{filepath.Join(serverAPIDir, "person", "person_api.go"), "person_api.txt"},
				{filepath.Join("types_lib", typeDir, "Person.go"), "types_lib/Person.txt"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}
		})

		Convey("client", t, func(c C) {
			var apiDef raml.APIDefinition

			err = raml.ParseFile("../fixtures/raml-examples/libraries/api.raml", &apiDef)
			c.So(err, ShouldBeNil)

			client, err := NewClient(&apiDef, "client", "examples.com/libro", targetDir, nil)
			c.So(err, ShouldBeNil)

			err = client.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "../fixtures/libraries/raml-examples/go_client"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"person_service.go", "person_service.txt"},
				{filepath.Join("types_lib", typeDir, "Person.go"), "types_lib/Person.txt"},
			}

			for _, check := range checks {
				s, err := utils.TestLoadFile(filepath.Join(targetDir, check.Result))
				c.So(err, ShouldBeNil)

				tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, check.Expected))
				c.So(err, ShouldBeNil)

				c.So(s, ShouldEqual, tmpl)
			}
		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})

}

func TestAliasLibTypeImportPath(t *testing.T) {
	Convey("TestAliasLibTypeImportPath", t, func(c C) {
		tests := []struct {
			path    string
			aliased string
		}{
			{"a.com/libraries/libname/types", `libname_types "a.com/libraries/libname/types"`},
			{"a.com/libname/types", `libname_types "a.com/libname/types"`},
		}

		for _, test := range tests {
			aliased := aliasLibTypeImportPath(test.path)
			c.So(aliased, ShouldEqual, test.aliased)
		}
	})
}
