package date

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeOnly(t *testing.T) {
	Convey("time-only", t, func(c C) {
		Convey("not in struct", t, func(c C) {
			dateStr := "10:09:08"

			// create time
			tim, err := time.Parse("15:04:05", dateStr)
			c.So(err, ShouldBeNil)

			to := TimeOnly(tim)

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
			jsonBytes := []byte(`{"name":"google","return":"10:09:08"}`)
			var data = struct {
				Name   string   `json:"name"`
				Return TimeOnly `json:"return"`
			}{}

			// unmarshal
			err := json.Unmarshal(jsonBytes, &data)
			c.So(err, ShouldBeNil)
			c.So(data.Return.String(), ShouldEqual, "10:09:08")

			// marshal again
			b, err := json.Marshal(&data)
			c.So(err, ShouldBeNil)
			c.So(string(b), ShouldEqual, string(jsonBytes))
		})
	})
}
