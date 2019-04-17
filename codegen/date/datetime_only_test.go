package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDatetimeOnly(t *testing.T) {
	Convey("datetime-only", t, func(c C) {
		Convey("not in struct", t, func(c C) {
			dateStr := "2015-07-04T21:00:00"

			// create time
			tim, err := time.Parse("2006-01-02T15:04:05.99", dateStr)
			c.So(err, ShouldBeNil)

			to := DatetimeOnly(tim)

			// marshal
			b, err := json.Marshal(&to)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, `"`+dateStr+`"`)

			// unmarshal
			err = json.Unmarshal([]byte(`"`+dateStr+`"`), &to)
			c.So(err, ShouldBeNil)
			c.So(to.String(), ShouldEqual, dateStr)
		})

		Convey("in struct", t, func(c C) {
			jsonBytes := []byte(`{"name":"google","born":"2015-07-04T21:00:00"}`)
			var data = struct {
				Name string       `json:"name"`
				Born DatetimeOnly `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			c.So(err, ShouldBeNil)
			c.So(data.Born.String(), ShouldEqual, "2015-07-04T21:00:00")

			// marshal again
			b, err := json.Marshal(&data)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
