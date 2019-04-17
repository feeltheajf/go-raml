package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestResourceTypeInheritance(t *testing.T) {
	apiDef := new(APIDefinition)
	err := ParseFile("./samples/resource_types.raml", apiDef)
	Convey("resource type & traits inheritance", t, func(c C) {
		c.So(err, ShouldBeNil)

		Convey("checking users", t, func(c C) {
			r := apiDef.Resources["/Users"]

			c.So(r.URI, ShouldEqual, "/Users")
			c.So(r.Description, ShouldEqual, "The collection of Users")

			c.So(r.Get, ShouldNotBeNil)
			c.So(r.Get.Description, ShouldEqual, "Get all Users, optionally filtered")
			c.So(r.Get.DisplayName, ShouldEqual, "ListAllUsers")
			c.So(r.Get.Responses["200"].Bodies.Type, ShouldEqual, "Users")

			c.So(r.Post, ShouldNotBeNil)
			c.So(r.Post.Description, ShouldEqual, "Create a new User")
			c.So(r.Post.Bodies.ApplicationJSON.Type, ShouldEqual, "User")
			c.So(r.Post.Responses["200"].Bodies.Type, ShouldEqual, "User")
		})

		Convey("checking queues - optional method", t, func(c C) {
			r := apiDef.Resources["/queues"]
			c.So(r, ShouldNotBeNil)

			c.So(r.Get, ShouldNotBeNil)
			c.So(r.Get.Description, ShouldEqual, "Get all queues")

			c.So(r.Post, ShouldBeNil)
		})

		Convey("checking corps - header - resourcePath - request body", t, func(c C) {
			r := apiDef.Resources["/corps"]
			c.So(r, ShouldNotBeNil)

			c.So(r.Post, ShouldNotBeNil)

			props := r.Post.Bodies.ApplicationJSON.Properties
			c.So(ToProperty("name", props["name"]).Type, ShouldEqual, "string")
			c.So(ToProperty("age", props["age"]).Type, ShouldEqual, "int")
			c.So(r.Post.Headers["X-Chargeback"].Required, ShouldBeTrue)

			mem := r.Nested["/{id}"]
			c.So(mem, ShouldNotBeNil)
			c.So(mem.Get.Description, ShouldEqual, "get /corps/{id}") // check resourcePath parsing

			// check resourcePathName parsing
			respCode := HTTPCode("200")
			c.So(mem.Get.Responses, ShouldContainKey, respCode)
			c.So(mem.Get.Responses[respCode].Bodies.Type, ShouldEqual, "corps")
		})

		Convey("books - query parameters", t, func(c C) {
			r := apiDef.Resources["/books"]
			c.So(r, ShouldNotBeNil)

			qps := r.Get.QueryParameters
			c.So(qps["title"].Description, ShouldEqual, "Return books that have their title matching the given value")
			c.So(qps["digest_all_fields"].Description, ShouldEqual,
				"If no values match the value given for title, use digest_all_fields instead")

			// collection merging
			// test disabled because of issue: https://github.com/Jumpscale/go-raml/issues/99
			//c.So(qps["platform"].Enum, ShouldContain, "mac")
			//c.So(qps["platform"].Enum, ShouldContain, "unix")
			//c.So(qps["platform"].Enum, ShouldContain, "win")
		})

		Convey("query parameters traits", t, func(c C) {
			r := apiDef.Resources["/books"]
			c.So(r, ShouldNotBeNil)

			c.So(apiDef.Traits, ShouldContainKey, "paged")
			c.So(r.Get, ShouldNotBeNil)

			qps := r.Get.QueryParameters
			numPages := qps["numPages"]
			c.So(numPages.Description, ShouldEqual, "The number of pages to return, not to exceed 10")
			c.So(numPages.Type, ShouldEqual, "integer")
			c.So(*numPages.Minimum, ShouldEqual, 1)
			c.So(numPages.Required, ShouldEqual, true)

			c.So(qps["access_token"].Description, ShouldEqual, "A valid access_token is required")

		})

		Convey("request body traits", t, func(c C) {
			r := apiDef.Resources["/servers"]
			c.So(r, ShouldNotBeNil)

			props := r.Post.Bodies.ApplicationJSON.Properties

			c.So(props, ShouldContainKey, "name")
			c.So(props, ShouldContainKey, "address?")
			c.So(props, ShouldNotContainKey, "location?")
			c.So(props, ShouldNotContainKey, "location")
		})
		Convey("resource types can use traits", t, func(c C) {
			c.So(apiDef.ResourceTypes, ShouldContainKey, "file")

			file := apiDef.ResourceTypes["file"]
			c.So(file.Put, ShouldNotBeNil)
			c.So(file.Put.Headers, ShouldContainKey, HTTPHeader("drm-key"))
			c.So(file.Put.Headers["drm-key"].Required, ShouldBeTrue)
		})
	})
}
