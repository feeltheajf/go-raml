package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateTimeRFC2616(t *testing.T) {

	Convey("datetime RF2616", t, func(c C) {
		Convey("not in struct", t, func(c C) {
			dateStr := "Sun, 28 Feb 2016 16:41:41 GMT"

			// create time
			tim, err := time.Parse(dateTimeRFC2616Fmt, dateStr)
			c.So(err, ShouldBeNil)

			to := DateTimeRFC2616(tim)

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
			jsonBytes := []byte(`{"name":"google","born":"Sun, 28 Feb 2016 16:41:41 GMT"}`)
			var data = struct {
				Name string          `json:"name"`
				Born DateTimeRFC2616 `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			c.So(err, ShouldBeNil)
			c.So(data.Born.String(), ShouldEqual, "Sun, 28 Feb 2016 16:41:41 GMT")

			// marshal again
			b, err := json.Marshal(&data)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, string(jsonBytes))
		})
	})

}
