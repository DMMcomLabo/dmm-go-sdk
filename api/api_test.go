package api

import (
	"reflect"
	"testing"
	// "github.com/davecgh/go-spew/spew"
)

const (
	Dummy_Affliate_Id = "foobar-999"
	Dummy_Api_Id      = "TXpEZ5D4T2xB3J5cuSLf"
)

func TestRequestJson(t *testing.T) {
	testUrl := "https://httpbin.org/get?foo=bar"
	actual, err1 := RequestJson(testUrl)

	if actual == nil {
		t.Fatalf("response is not expected empty.")
	}

	if err1 != nil {
		t.Fatalf("error is expected empty.")
	}

	if reflect.TypeOf(actual).String() != "map[string]interface {}" {
		t.Fatalf("response is expected inteface{}. but actual %s", reflect.TypeOf(actual).String())
	}

	errUrl := "https://httpbin.org/status/500"
	_, err2 := RequestJson(errUrl)

	if err2 == nil {
		t.Fatalf("error is not expected empty.")
	}
}

func TestValidateAffiliateId(t *testing.T) {
	val := "vcder56yuhnmkiuy-990"
	if !ValidateAffiliateId(val) {
		t.Fatalf("When value is %s, not expected false.", val)
	}

	val = "vcder56yuhnmkiuy-9"
	if ValidateAffiliateId(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	val = "vcder56yuhnmkiuy-"
	if ValidateAffiliateId(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	val = "-999"
	if ValidateAffiliateId(val) {
		t.Fatalf("When value is %s, not expected true.", val)
	}

	if ValidateAffiliateId("") {
		t.Fatalf("When value is empty, not expected true.")
	}
}

func TestValidateSite(t *testing.T) {
	if !ValidateSite(SITE_ALLAGES) {
		t.Fatalf("When value is %s, not expected false.", SITE_ALLAGES)
	}

	if !ValidateSite(SITE_ADULT) {
		t.Fatalf("When value is %s, not expected false.", SITE_ADULT)
	}

	if !ValidateSite("DMM.com") {
		t.Fatalf("When value is %s, not expected false.", "DMM.com")
	}

	if !ValidateSite("DMM.R18") {
		t.Fatalf("When value is %s, not expected false.", "DMM.R18")
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
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
	}

	target = 1
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
	}

	target = 100
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
	}

	target = 0
	if ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected true.", target, min, max)
	}

	target = 101
	if ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected true.", target, min, max)
	}

	target = 10
	min = 10
	max = 10
	if !ValidateRange(target, min, max) {
		t.Fatalf("When target value is %d, min is %d and max is %d, not expected false.", target, min, max)
	}
}

func TestGetApiVersion(t *testing.T) {
	if GetApiVersion() != API_VERSION {
		t.Errorf("This value is expected to equal API_VERSION. actual:%s", GetApiVersion())
	}
}
