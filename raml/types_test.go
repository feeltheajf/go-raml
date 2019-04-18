package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeInType(t *testing.T) {
	apiDef := new(APIDefinition)
	Convey("Type in type's properties", t, func(c C) {
		err := ParseFile("./samples/types.raml", apiDef)
		c.So(err, ShouldBeNil)

		action, ok := apiDef.Types["Action"]
		c.So(ok, ShouldBeTrue)

		name := action.GetProperty("name")
		c.So(name.Type, ShouldEqual, "string")

		recurring := action.GetProperty("recurring")
		c.So(recurring.TypeString(), ShouldEqual, "Actionrecurring")

		// check the inline type
		ar, ok := apiDef.Types["Actionrecurring"]
		c.So(ok, ShouldBeTrue)

		// Must work via .GetPropert
		period := ar.GetProperty("period")
		c.So(period.TypeString(), ShouldEqual, "integer")

		// Also must work via .ToProperty
		var prop Property
		for k, p := range action.Properties {
			if k == "recurring" {
				prop = ToProperty(k, p)
				break
			}
		}
		c.So(prop.TypeString(), ShouldEqual, "Actionrecurring")

		// test for the recursive type
		_, ok = apiDef.Types["Actionrecurringcombo"]
		c.So(ok, ShouldBeTrue)

		combo := ar.GetProperty("combo")
		c.So(combo.TypeString(), ShouldEqual, "Actionrecurringcombo")

		// check the items with properties
		coinInputs := action.GetProperty("coininputs")
		c.So(coinInputs.Type, ShouldEqual, "array")
		c.So(coinInputs.Items.Type, ShouldEqual, "ActioncoininputsItem")

		_, ok = apiDef.Types["ActioncoininputsItem"]
		c.So(ok, ShouldBeTrue)

		// check the items with Type and format
		coinTipes := action.GetProperty("coinTipes")
		c.So(coinTipes.Type, ShouldEqual, "array")
		c.So(coinTipes.Items.Type, ShouldEqual, "number")
		c.So(coinTipes.Items.Format, ShouldEqual, "double")

		// check the items with plain type
		coinTipesPlain := action.GetProperty("coinTipesPlain")
		c.So(coinTipesPlain.Type, ShouldEqual, "array")
		c.So(coinTipesPlain.Items.Type, ShouldEqual, "string")
	})
}
