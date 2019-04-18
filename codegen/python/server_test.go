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

func TestServer(t *testing.T) {
	Convey("server generator", t, func(c C) {
		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		apiDef := new(raml.APIDefinition)
		err = raml.ParseFile("../fixtures/congo/api.raml", apiDef)
		c.So(err, ShouldBeNil)
		rootFixture := "../fixtures/congo/python_server"
		files := []string{
			"drones_api.py",
			"deliveries_api.py",
			"app.py",
			"handlers/__init__.py",
			"handlers/drones_postHandler.py",
		}

		Convey("Congo python server", t, func(c C) {
			server := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, false)
			err = server.Generate()
			c.So(err, ShouldBeNil)
			validateFiles(c, files, targetDir, rootFixture)
			// test that this file exist
			filesExist := []string{
				"types/User.py",
				"types/client_support.py",
				"handlers/deliveries_getHandler.py",
				"handlers/deliveries_postHandler.py",
				"handlers/deliveries_deliveryId_getHandler.py",
				"handlers/deliveries_deliveryId_patchHandler.py",
				"handlers/deliveries_deliveryId_deleteHandler.py",
				"handlers/drones_getHandler.py",
				"handlers/drones_postHandler.py",
				"handlers/drones_droneId_getHandler.py",
				"handlers/drones_droneId_patchHandler.py",
				"handlers/drones_droneId_deleteHandler.py",
				"handlers/drones_droneId_deliveries_getHandler.py",
			}
			for _, f := range filesExist {
				_, err := os.Stat(filepath.Join(targetDir, f))
				c.So(err, ShouldBeNil)
			}
		})

		Convey("Congo gevent python server", t, func(c C) {
			files = append(files, "server.py")
			server := NewFlaskServer(apiDef, "apidocs", targetDir, true, nil, true)
			err = server.Generate()
			c.So(err, ShouldBeNil)
			validateFiles(c, files, targetDir, rootFixture)
		})

		c.Reset(func() {
			os.RemoveAll(targetDir)
		})
	})
}

func validateFiles(c C, files []string, targetDir string, rootFixture string) {
	for _, f := range files {
		s, err := utils.TestLoadFile(filepath.Join(targetDir, f))
		c.So(err, ShouldBeNil)

		tmpl, err := utils.TestLoadFile(filepath.Join(rootFixture, f))
		c.So(err, ShouldBeNil)

		c.So(s, ShouldEqual, tmpl)
	}
}
