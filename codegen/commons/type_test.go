package commons

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/feeltheajf/go-raml/raml"
)

func TestCheckDuplicatedTitleTypes(t *testing.T) {
	Convey("TestCheckDuplicatedTitleTypes", t, func(c C) {
		tests := []struct {
			apiDef *raml.APIDefinition
			err    bool
		}{
			{
				&raml.APIDefinition{
					Types: map[string]raml.Type{
						"One": raml.Type{},
						"one": raml.Type{},
					},
				}, true,
			},
			{
				&raml.APIDefinition{
					Types: map[string]raml.Type{
						"One": raml.Type{},
						"oNe": raml.Type{},
					},
				}, false,
			},
		}

		for _, test := range tests {
			err := CheckDuplicatedTitleTypes(test.apiDef)
			if test.err {
				c.So(err, ShouldNotBeNil)
			} else {
				c.So(err, ShouldBeNil)
			}
		}

	})
}
