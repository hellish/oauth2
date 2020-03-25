package manage_test

import (
	"testing"

	"gopkg.in/hellish/oauth2.v3/manage"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUtil(t *testing.T) {
	Convey("Util Test", t, func() {
		Convey("ValidateURI Test", func() {
			err := manage.DefaultValidateURI("http://www.example.com", "http://www.example.com/cb?code=xxx")
			So(err, ShouldBeNil)
		})
	})
	Convey("DefaultValidateClientSecret Test", t, func() {
		result1 := manage.DefaultValidateClientSecret("password1", "password1")
		So(result1, ShouldBeTrue)
		result2 := manage.DefaultValidateClientSecret("password1", "password2")
		So(result2, ShouldBeFalse)
	})
}
