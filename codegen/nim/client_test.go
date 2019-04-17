package nim

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Jumpscale/go-raml/raml"
	"github.com/Jumpscale/go-raml/utils"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateClient(t *testing.T) {
	Convey("generate client from raml", t, func(c C) {
		var apiDef raml.APIDefinition
		err := raml.ParseFile("../fixtures/client_resources/client.raml", &apiDef)
		c.So(err, ShouldBeNil)

		targetDir, err := ioutil.TempDir("", "")
		c.So(err, ShouldBeNil)

		client := Client{
			APIDef: &apiDef,
			Dir:    targetDir,
		}
		err = client.Generate()
		c.So(err, ShouldBeNil)

		rootFixture := "./fixtures/resource/client"
		checks := []struct {
			Result   string
			Expected string
		}{
			{"client_struct.nim", "client_struct.nim"},
			{"Users_service.nim", "Users_service.nim"},
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
