package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLibraries(t *testing.T) {
	Convey("Libraries", t, func(c C) {
		apiDef := new(APIDefinition)
		err := ParseFile("./samples/simple_with_lib.raml", apiDef)
		c.So(err, ShouldBeNil)

		Convey("two level library", t, func(c C) {
			// check Uses
			c.So(apiDef.Uses, ShouldContainKey, "files")
			c.So(apiDef.Uses["files"], ShouldEqual, "libraries/files.raml")

			// Check Libraries property
			c.So(apiDef.Libraries, ShouldContainKey, "files")

			// first level
			files := apiDef.Libraries["files"]
			c.So(files.Usage, ShouldEqual, "Use to define some basic file-related constructs.")
			c.So(files.Traits, ShouldContainKey, "drm")
			c.So(files.Uses, ShouldContainKey, "file-type")
			c.So(files.ResourceTypes, ShouldContainKey, "file")

			// check trait usage in a resource type
			file := files.ResourceTypes["file"]
			c.So(file.Get, ShouldNotBeNil)
			c.So(file.Get.Headers, ShouldContainKey, HTTPHeader("drm-key"))

			// second level
			c.So(files.Libraries, ShouldContainKey, "file-type")
			fileType := files.Libraries["file-type"]
			c.So(fileType.Types, ShouldContainKey, "File")
			File := fileType.Types["File"]
			c.So(len(File.Properties), ShouldEqual, 2)
		})

		Convey("using library's trait in root's definition", t, func(c C) {
			files := apiDef.Resources["/files"]
			c.So(files.Get, ShouldNotBeNil)
			c.So(files.Get.Headers, ShouldContainKey, HTTPHeader("drm-key"))
			c.So(files.Get.Headers["drm-key"].Required, ShouldBeFalse)
		})

		Convey("proper variable name", t, func(c C) {
			r := apiDef.Resources["/links"]

			c.So(r.Post, ShouldNotBeNil)
			c.So(r.Post.Bodies.Type, ShouldEqual, "files.Link")
		})

	})
}
