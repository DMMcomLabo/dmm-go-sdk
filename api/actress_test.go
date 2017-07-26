package api

import (
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestNewActressService(t *testing.T) {
	affiliateID := DummyAffliateID
	apiID := DummyAPIID

	srv := NewActressService(affiliateID, apiID)
	if srv.AffiliateID != affiliateID {
		t.Fatalf("ActressService.AffiliateID is expected to equal the input value(affiliateID)")
	}

	if srv.ApiID != apiID {
		t.Fatalf("ActressService.ApiID is expected to equal the input value(apiID)")
	}
}

func TestSetLengthInActressService(t *testing.T) {
	srv := dummyActressService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("ActressService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInActressService(t *testing.T) {
	srv := dummyActressService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("ActressService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInActressService(t *testing.T) {
	srv := dummyActressService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("ActressService.Offset is expected to equal the input value(offset)")
	}
}

func TestSetKeywordInActressService(t *testing.T) {
	srv := dummyActressService()

	keyword1 := "abcdefghijkelmnopqrstuvwxyzABCDEFGHIJKELMNOPQRSTUVWXYZ0123456789"
	srv.SetKeyword(keyword1)
	if srv.Keyword != keyword1 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword1)")
	}

	keyword2 := ""
	srv.SetKeyword(keyword2)
	if srv.Keyword != keyword2 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword2)")
	}

	keyword3 := "つれづれなるまゝに、日暮らし、硯にむかひて、心にうつりゆくよしなし事を、そこはかとなく書きつくれば、あやしうこそものぐるほしけれ。"
	srv.SetKeyword(keyword3)
	if srv.Keyword != keyword3 {
		t.Fatalf("ActressService.Keyword is expected to equal the input value(keyword3)")
	}

	keyword4 := " a b c d 0 "
	keyword4Expected := "a b c d 0"
	srv.SetKeyword(keyword4)
	if srv.Keyword != keyword4Expected {
		t.Fatalf("ActressService.Keyword is expected to equal keyword4_expected")
	}

	keyword5 := "　あ ア　化Ａ "
	keyword5Expected := "あ ア　化Ａ"
	srv.SetKeyword(keyword5)
	if srv.Keyword != keyword5Expected {
		t.Fatalf("ActressService.Keyword is expected to equal keyword5_expected")
	}
}

func TestSetBirthdayInActressService(t *testing.T) {
	srv := dummyActressService()

	date1 := "19840723"
	srv.SetBirthday(date1)
	if srv.Birthday != date1 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date1)")
	}

	date2 := ""
	srv.SetBirthday(date2)
	if srv.Birthday != date2 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date2)")
	}

	date3 := "19000101-20160101"
	srv.SetBirthday(date3)
	if srv.Birthday != date3 {
		t.Fatalf("ActressService.Birthday is expected to equal the input value(date3)")
	}
}

func TestSetBustInActressService(t *testing.T) {
	srv := dummyActressService()

	bust1 := "D"
	srv.SetBust(bust1)
	if srv.Bust != bust1 {
		t.Fatalf("ActressService.Bust is expected to equal the input value(bust1)")
	}

	bust2 := ""
	srv.SetBust(bust2)
	if srv.Bust != bust2 {
		t.Fatalf("ActressService.Bust is expected to equal the input value(bust2)")
	}
}

func TestSetWaistInActressService(t *testing.T) {
	srv := dummyActressService()

	waist1 := "60"
	srv.SetWaist(waist1)
	if srv.Waist != waist1 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist1)")
	}

	waist2 := ""
	srv.SetWaist(waist2)
	if srv.Waist != waist2 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist2)")
	}

	waist3 := "50-60"
	srv.SetWaist(waist3)
	if srv.Waist != waist3 {
		t.Fatalf("ActressService.Waist is expected to equal the input value(waist3)")
	}
}

func TestSetHipInActressService(t *testing.T) {
	srv := dummyActressService()

	hip1 := "88"
	srv.SetHip(hip1)
	if srv.Hip != hip1 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip1)")
	}

	hip2 := ""
	srv.SetHip(hip2)
	if srv.Hip != hip2 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip2)")
	}

	hip3 := "80-90"
	srv.SetHip(hip3)
	if srv.Hip != hip3 {
		t.Fatalf("ActressService.Hip is expected to equal the input value(hip3)")
	}
}

func TestSetHeightInActressService(t *testing.T) {
	srv := dummyActressService()

	height1 := "155"
	srv.SetHeight(height1)
	if srv.Height != height1 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height1)")
	}

	height2 := ""
	srv.SetHeight(height2)
	if srv.Height != height2 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height2)")
	}

	height3 := "140-160"
	srv.SetHeight(height3)
	if srv.Height != height3 {
		t.Fatalf("ActressService.Height is expected to equal the input value(height3)")
	}
}

func TestValidateLengthInActressService(t *testing.T) {
	srv := dummyActressService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DefaultActressAPILength
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DefaultActressMaxLength
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ActressService.ValidateLength is expected TRUE.")
	}

	target = DefaultActressMaxLength + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ActressService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInActressService(t *testing.T) {
	srv := dummyActressService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ActressService.ValidateOffset is expected TRUE.")
	}

	target = 100
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ActressService.ValidateOffset is expected TRUE.")
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ActressService.ValidateOffset is expected FALSE.")
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ActressService.ValidateOffset is expected FALSE.")
	}
}

func TestBuildRequestURLInActressService(t *testing.T) {
	var srv *ActressService
	var u string
	var err error
	var expected string

	srv = dummyActressService()
	u, err = srv.BuildRequestURL()
	expected = APIBaseURL + "/ActressSearch?affiliate_id=" + DummyAffliateID + "&api_id=" + DummyAPIID + "&hits=" + strconv.FormatInt(DefaultActressAPILength, 10) + "&offset=" + strconv.FormatInt(DefaultAPIOffset, 10)
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}

	srv = dummyActressService()
	srv.SetLength(0)
	srv.SetOffset(0)
	u, err = srv.BuildRequestURL()
	expected = APIBaseURL + "/ActressSearch?affiliate_id=" + DummyAffliateID + "&api_id=" + DummyAPIID
	expectedBase := expected
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}

	srv.SetLength(-1)
	u, err = srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestURL is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestURL is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestURL is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestURL is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetInitial("あ")
	expected = expectedBase + "&initial=" + url.QueryEscape("あ")
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetInitial("")

	srv.SetSort("name")
	expected = expectedBase + "&sort=name"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetSort("")

	srv.SetKeyword("天使もえ")
	expected = expectedBase + "&keyword=" + url.QueryEscape("天使もえ")
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetKeyword("")

	srv.SetBirthday("19940710")
	expected = expectedBase + "&birthday=19940710"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetBirthday("")

	srv.SetGteBirthday("19800201")
	expected = expectedBase + "&gte_birthday=19800201"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetGteBirthday("")

	srv.SetLteBirthday("19940710")
	expected = expectedBase + "&lte_birthday=19940710"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetLteBirthday("")

	srv.SetBust("84")
	expected = expectedBase + "&bust=84"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetBust("")

	srv.SetGteBust("1")
	expected = expectedBase + "&gte_bust=1"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetGteBust("")

	srv.SetLteBust("100")
	expected = expectedBase + "&lte_bust=100"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetLteBust("")

	srv.SetWaist("57")
	expected = expectedBase + "&waist=57"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetWaist("")

	srv.SetGteWaist("51")
	expected = expectedBase + "&gte_waist=51"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetGteWaist("")

	srv.SetLteWaist("60")
	expected = expectedBase + "&lte_waist=60"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetLteWaist("")

	srv.SetHip("82")
	expected = expectedBase + "&hip=82"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetHip("")

	srv.SetGteHip("80")
	expected = expectedBase + "&gte_hip=80"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetGteHip("")

	srv.SetLteHip("90")
	expected = expectedBase + "&lte_hip=90"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetLteHip("")

	srv.SetHeight("155")
	expected = expectedBase + "&height=155"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetHeight("")

	srv.SetGteHeight("150")
	expected = expectedBase + "&gte_height=150"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetGteHeight("")

	srv.SetLteHeight("160")
	expected = expectedBase + "&lte_height=160"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetLteHeight("")
}

func TestBuildRequestURLWithoutApiIDInActressService(t *testing.T) {
	srv := dummyActressService()
	srv.ApiID = ""
	u, err := srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestURL is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestURL is expected to return error.")
	}
}

func TestBuildRequestURLWithWrongAffiliateIDInActressService(t *testing.T) {
	srv := dummyActressService()
	srv.AffiliateID = "fizzbizz-100"
	u, err := srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ActressService.BuildRequestURL is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("ActressService.BuildRequestURL is expected to return error.")
	}
}

func TestExcuteWeakRequestActressAPIToServer(t *testing.T) {
	if !RequestAvailable {
		t.Skip("Not set valid credentials")
	}

	srv := NewActressService(TestAffiliateID, TestAPIID)
	srv.SetInitial("あ")
	srv.SetKeyword("あさみ")
	srv.SetBust("90")
	srv.SetGteBust("90")
	srv.SetLteBust("99")
	srv.SetWaist("-60")
	srv.SetGteWaist("50")
	srv.SetLteWaist("60")
	srv.SetHip("85-90")
	srv.SetGteHip("85")
	srv.SetLteHip("90")
	srv.SetHeight("160")
	srv.SetGteHeight("160")
	srv.SetLteHeight("169")
	srv.SetBirthday("19900101")
	srv.SetGteBirthday("19900101")
	srv.SetLteBirthday("20101231")
	srv.SetSort("-name")
	srv.SetLength(20)
	srv.SetOffset(1)

	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, The network is down.")
	}

	if reflect.TypeOf(rst).String() != "*api.ActressResponse" {
		t.Fatalf("ActressService.Execute is expected to return *api.ActressResponse")
	}
}

func dummyActressService() *ActressService {
	return NewActressService(DummyAffliateID, DummyAPIID)
}
