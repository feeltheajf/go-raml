package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateOnly(t *testing.T) {
	Convey("date-only", t, func(c C) {
		Convey("not in struct", t, func(c C) {
			dateStr := "2016-05-04"

			// create time
			tim, err := time.Parse("2006-01-02", dateStr)
			c.So(err, ShouldBeNil)

			do := DateOnly(tim)

			// marshal
			b, err := json.Marshal(&do)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, `"`+dateStr+`"`)

			// unmarshal
			err = json.Unmarshal([]byte(`"`+dateStr+`"`), &do)
			c.So(err, ShouldBeNil)
			c.So(do.String(), ShouldEqual, dateStr)
		})

		Convey("in struct", t, func(c C) {
			jsonBytes := []byte(`{"name":"google","born":"2016-05-04"}`)
			var data = struct {
				Name string   `json:"name"`
				Born DateOnly `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			c.So(err, ShouldBeNil)

			// marshal again
			b, err := json.Marshal(&data)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
