package langreg

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegionCodeInfo(t *testing.T) {
	var tests = []struct {
		code            string
		expectedName    string
		expectedccTLD   string
		errExpected     bool
		testDescription string
	}{
		{"GB", "United Kingdom", ".uk", false, "code with ascii character set"},
		{"gb", "United Kingdom", ".uk", false, "code with lowercase ascii characters"},
		{"AX", "Aland Islands !Åland Islands", ".ax", false, "code with non-ascii Unicode character set"},
		{" GB ", "United Kingdom", ".uk", false, "code with leading and trailng space"},
		{"\n\tGB", "United Kingdom", ".uk", false, "code with leading and trailing whitespace"},
		{"RSA", "", "", true, "code that is too long"},
		{"Z", "", "", true, "code that is too short"},
		{"ZZ", "", "", true, "code that is invalid"},
		{"*G", "", "", true, "code with invalid characters"},
		{"73", "", "", true, "code with numbers"},
	}
	for _, tt := range tests {
		Convey("Given a "+tt.testDescription, t, func() {
			actualName, actualCCTLD, actualErr := RegionCodeInfo(tt.code)
			So(actualName, ShouldEqual, tt.expectedName)
			So(actualCCTLD, ShouldEqual, tt.expectedccTLD)
			if !tt.errExpected {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldNotBeNil)
			}
		})
	}
}
