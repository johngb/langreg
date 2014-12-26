package langreg

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLanguageCodeInfo(t *testing.T) {
	var tests = []struct {
		code            string
		expectedEn      string
		expectedNat     string
		errExpected     bool
		testDescription string
	}{
		{"en", "English", "English", false, "code with ascii character set"},
		{"EN", "English", "English", false, "code with uppercase ascii characters"},
		{"iu", "Inuktitut", "ᐃᓄᒃᑎᑐᑦ", false, "code with non-ascii Unicode character set"},
		{"  en ", "English", "English", false, "code with leading and trailng space"},
		{"\nen ", "English", "English", false, "code with leading and trailing whitespace"},
		{"zzz", "", "", true, "code that is too long"},
		{"z", "", "", true, "code that is too short"},
		{"zz", "", "", true, "code that is invalid"},
		{"*z", "", "", true, "code with invalid characters"},
		{"73", "", "", true, "code with numbers"},
	}
	for _, tt := range tests {
		Convey("Given a "+tt.testDescription, t, func() {
			actualEn, actualNat, actualErr := LangCodeInfo(tt.code)
			So(actualEn, ShouldEqual, tt.expectedEn)
			So(actualNat, ShouldEqual, tt.expectedNat)
			if !tt.errExpected {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldNotBeNil)
			}
		})
	}
}
