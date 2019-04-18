package golang

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetImportPath(t *testing.T) {
	Convey("TestSetImportPath", t, func(c C) {
		oriGoPath := os.Getenv("GOPATH")
		Convey("users api", t, func(c C) {
			fakeGopath := "/gopath"
			os.Setenv("GOPATH", fakeGopath)
			c.So(setRootImportPath("import.com/a", "target"), ShouldEqual, "import.com/a")
			c.So(setRootImportPath("", "/gopath/src/johndoe.com/cool"), ShouldEqual, "johndoe.com/cool")
		})

		c.Reset(func() {
			os.Setenv("GOPATH", oriGoPath)
		})
	})
}
