package types

import (
	"testing"

	"github.com/feeltheajf/go-raml/raml"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateStructBodyFromRaml(t *testing.T) {
	Convey("generate struct body from raml", t, func(c C) {
		apiDef := new(raml.APIDefinition)

		Convey("simple raml", t, func() {
			err := raml.ParseFile("./fixtures/api.raml", apiDef)
			c.So(err, ShouldBeNil)

			expected := []string{
				"/users:POST:body:0",
				"/users/{id}:GET:body:200",
				"[Cat,animal]",
				"Cat | animal",
				"animal",
				"Cat",
				"EnumCity",
			}
			tts := AllTypes(apiDef, "main")

			c.So(len(tts), ShouldEqual, len(expected))
			for name := range tts {
				c.So(expected, ShouldContain, name)
				t.Logf("name=%v\n", name)
			}
		})
	})
}
