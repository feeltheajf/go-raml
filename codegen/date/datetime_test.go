package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateTime(t *testing.T) {
	Convey("datetime RFC3339", t, func(c C) {
		Convey("not in struct", t, func(c C) {
			dateStr := "2016-02-28T16:41:41.09Z"

			// create time
			tim, err := time.Parse(dateTimeFmt, dateStr)
			c.So(err, ShouldBeNil)

			to := DateTime(tim)

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
			jsonBytes := []byte(`{"name":"google","born":"2016-02-28T16:41:41.09Z"}`)
			var data = struct {
				Name string   `json:"name"`
				Born DateTime `json:"born"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			c.So(err, ShouldBeNil)
			c.So(data.Born.String(), ShouldEqual, "2016-02-28T16:41:41.09Z")

			// marshal again
			b, err := json.Marshal(&data)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, string(jsonBytes))
		})
	})

}
