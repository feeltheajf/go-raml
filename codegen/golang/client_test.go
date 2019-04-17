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

func TestClientBasic(t *testing.T) {
	Convey("Test client", t, func(c C) {
		var apiDef raml.APIDefinition
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("users api", t, func(c C) {
			err = raml.ParseFile("../fixtures/client_resources/client.raml", &apiDef)
			c.So(err, ShouldBeNil)

			client, err := NewClient(&apiDef, "theclient", "examples.com/libro", targetDir, nil)
			c.So(err, ShouldBeNil)

			err = client.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/client_resources"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"users_service.go", "users_service.txt"},
				{"client_structapitest.go", "client_structapitest.txt"},
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

func TestClientMultiSlashEndpoint(t *testing.T) {
	Convey("Test client with multislash endpoint", t, func(c C) {
		var apiDef raml.APIDefinition
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		Convey("users api", t, func(c C) {
			err = raml.ParseFile("../fixtures/client_resources/multislash.raml", &apiDef)
			c.So(err, ShouldBeNil)

			client, err := NewClient(&apiDef, "theclient", "examples.com/libro", targetDir, nil)
			c.So(err, ShouldBeNil)

			err = client.Generate()
			c.So(err, ShouldBeNil)

			rootFixture := "./fixtures/client_resources/multislash"
			checks := []struct {
				Result   string
				Expected string
			}{
				{"client_goramldir.go", "client_goramldir.txt"},
				{"animalsid_service.go", "animalsid_service.txt"},
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
