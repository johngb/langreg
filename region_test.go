package langreg

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

type regTestInfo struct {
	code            string
	expectedName    string
	expectedccTLD   string
	errExpected     bool
	testDescription string
}

var regTests = []regTestInfo{
	{"GB", "United Kingdom", ".uk", false, "code with ascii character set"},
	{"gb", "", "", true, "code with lowercase ascii characters"},
	{"AX", "Aland Islands !Åland Islands", ".ax", false, "code with non-ascii Unicode character set"},
	{" GB ", "", "", true, "code with leading and trailng space"},
	{"\n\tGB", "", "", true, "code with leading and trailing whitespace"},
	{"RSA", "", "", true, "code that is too long"},
	{"Z", "", "", true, "code that is too short"},
	{"ZZ", "", "", true, "code that is invalid"},
	{"*G", "", "", true, "code with invalid characters"},
	{"73", "", "", true, "code with numbers"},
}

func TestRegionCodeInfo(t *testing.T) {
	for _, tt := range regTests {
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

func TestIsValidRegionCode(t *testing.T) {
	for _, tt := range regTests {
		Convey("Given a "+tt.testDescription, t, func() {
			actualIsValid := IsValidRegionCode(tt.code)
			So(!actualIsValid, ShouldEqual, tt.errExpected)
		})
	}
}

func TestRegionName(t *testing.T) {
	for _, tt := range regTests {
		Convey("Given a "+tt.testDescription, t, func() {
			actualName, actualErr := RegionName(tt.code)
			So(actualName, ShouldEqual, tt.expectedName)
			if !tt.errExpected {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldNotBeNil)
			}
		})
	}
}

func TestRegionCCTLD(t *testing.T) {
	for _, tt := range regTests {
		Convey("Given a "+tt.testDescription, t, func() {
			actualCCTLD, actualErr := RegionCCTLD(tt.code)
			So(actualCCTLD, ShouldEqual, tt.expectedccTLD)
			if !tt.errExpected {
				So(actualErr, ShouldBeNil)
			} else {
				So(actualErr, ShouldNotBeNil)
			}
		})
	}
}
