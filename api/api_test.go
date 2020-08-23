package api

import (
	"os"
	"reflect"
	"testing"
)

const (
	DummyAffliateID = "foobar-999"
	DummyAPIID      = "TXpEZ5D4T2xB3J5cuSLf"
)

var (
	TestAffiliateID  string
	TestAPIID        string
	RequestAvailable bool
)

func init() {
	TestAffiliateID = os.Getenv("DMM_TEST_AFFILIATE_ID")
	TestAPIID = os.Getenv("DMM_TEST_API_ID")

	RequestAvailable = (ValidateAffiliateID(TestAffiliateID) && TestAPIID != "")
}

func TestRequestJSON(t *testing.T) {
	testURL := "https://httpbin.org/get?foo=bar"
	actual, err1 := RequestJSON(testURL)

	if actual == nil {
		t.Fatalf("response is not expected empty.")
	}

	if err1 != nil {
		t.Fatalf("error is expected empty.")
	}

	if reflect.TypeOf(actual).String() != "map[string]interface {}" {
		t.Fatalf("response is expected inteface{}. but actual %s", reflect.TypeOf(actual).String())
	}

	errURL := "https://httpbin.org/status/500"
	_, err2 := RequestJSON(errURL)

	if err2 == nil {
		t.Fatalf("error is not expected empty.")
	}
}

func TestValidateAffiliateID(t *testing.T) {
	val := "vcder56yuhnmkiuy-990"
	if !ValidateAffiliateID(val) {
		t.Fatalf("When value is %s, not expected false.", val)
	}

	val = "vcder56yuhnmkiuy-9"
	if ValidateAffiliateID(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	val = "vcder56yuhnmkiuy-"
	if ValidateAffiliateID(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	val = "-999"
	if ValidateAffiliateID(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	if ValidateAffiliateID("") {
		t.Fatalf("When value is empty, not expected true.")
	}
}

func TestValidateSite(t *testing.T) {
	if !ValidateSite(SiteGeneral) {
		t.Fatalf("When value is %s, not expected false.", SiteGeneral)
	}

	if !ValidateSite(SiteAdult) {
		t.Fatalf("When value is %s, not expected false.", SiteAdult)
	}

	if ValidateSite("DMM.co.jp") {
		t.Fatalf("When value is %s, not expected true.", "DMM.co.jp")
	}

	if ValidateSite("") {
		t.Fatalf("When value is empty, not expected true.")
	}
}

func TestTrimString(t *testing.T) {
	expected1 := "abcdefghijkelmnopqrstuvwxyzABCDEFGHIJKELMNOPQRSTUVWXYZ0123456789"
	actual1 := TrimString(expected1)
	if expected1 != actual1 {
		t.Fatalf("trimed string is expected to equal expected1")
	}

	expected2 := ""
	actual2 := TrimString(expected2)
	if expected2 != actual2 {
		t.Fatalf("trimed string is expected to equal expected2")
	}

	expected3 := "つれづれなるまゝに、日暮らし、硯にむかひて、心にうつりゆくよしなし事を、そこはかとなく書きつくれば、あやしうこそものぐるほしけれ。"
	actual3 := TrimString(expected3)
	if expected3 != actual3 {
		t.Fatalf("trimed string is expected to equal expected3")
	}

	target4 := " a b c d 0 "
	expected4 := "a b c d 0"
	actual4 := TrimString(target4)
	if expected4 != actual4 {
		t.Fatalf("trimed string is expected to equal expected4")
	}

	target5 := "　あ ア　化Ａ "
	expected5 := "あ ア　化Ａ"
	actual5 := TrimString(target5)
	if expected5 != actual5 {
		t.Fatalf("trimed string is expected to equal expected5")
	}
}

func TestValidateRange(t *testing.T) {
	var target int64
	var min int64
	var max int64

	target = 10
	min = 1
	max = 100
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false", target, min, max)
	}

	target = 1
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false", target, min, max)
	}

	target = 100
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false", target, min, max)
	}

	target = 0
	if ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected true", target, min, max)
	}

	target = 101
	if ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected true", target, min, max)
	}

	target = 10
	min = 10
	max = 10
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false", target, min, max)
	}
}

func TestGetAPIVersion(t *testing.T) {
	if GetAPIVersion() != APIVersion {
		t.Errorf("This value is expected to equal APIVersion. actual:%s", GetAPIVersion())
	}
}
