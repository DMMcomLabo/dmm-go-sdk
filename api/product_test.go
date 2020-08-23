package api

import (
	"net/url"
	"reflect"
	"strconv"
	"testing"
)

func TestNewProductService(t *testing.T) {
	affiliateID := DummyAffliateID
	apiID := DummyAPIID

	srv := NewProductService(affiliateID, apiID)
	if srv.AffiliateID != affiliateID {
		t.Fatalf("ProductService.AffiliateID is expected to equal the input value(affiliateID)")
	}

	if srv.ApiID != apiID {
		t.Fatalf("ProductService.ApiID is expected to equal the input value(apiID)")
	}
}

func TestSetLengthInProductService(t *testing.T) {
	srv := dummyProductService()
	var length int64 = 10
	srv.SetLength(length)

	if srv.Length != length {
		t.Fatalf("ProductService.Length is expected to equal the input value(length)")
	}
}

func TestSetHitsInProductService(t *testing.T) {
	srv := dummyProductService()
	var hits int64 = 10
	srv.SetHits(hits)

	if srv.Length != hits {
		t.Fatalf("ProductService.Length is expected to equal the input value(hits)")
	}
}

func TestSetOffsetInProductService(t *testing.T) {
	srv := dummyProductService()
	var offset int64 = 10
	srv.SetOffset(offset)

	if srv.Offset != offset {
		t.Fatalf("ProductService.Offset is expected to equal the input value(offset)")
	}
}

func TestSetKeywordInProductService(t *testing.T) {
	srv := dummyProductService()

	keyword1 := "abcdefghijkelmnopqrstuvwxyzABCDEFGHIJKELMNOPQRSTUVWXYZ0123456789"
	srv.SetKeyword(keyword1)
	if srv.Keyword != keyword1 {
		t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword1)")
	}

	keyword2 := ""
	srv.SetKeyword(keyword2)
	if srv.Keyword != keyword2 {
		t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword2)")
	}

	keyword3 := "つれづれなるまゝに、日暮らし、硯にむかひて、心にうつりゆくよしなし事を、そこはかとなく書きつくれば、あやしうこそものぐるほしけれ。"
	srv.SetKeyword(keyword3)
	if srv.Keyword != keyword3 {
		t.Fatalf("ProductService.Keyword is expected to equal the input value(keyword3)")
	}

	keyword4 := " a b c d 0 "
	keyword4Expected := "a b c d 0"
	srv.SetKeyword(keyword4)
	if srv.Keyword != keyword4Expected {
		t.Fatalf("ProductService.Keyword is expected to equal keyword4_expected")
	}

	keyword5 := "　あ ア　化Ａ "
	keyword5Expected := "あ ア　化Ａ"
	srv.SetKeyword(keyword5)
	if srv.Keyword != keyword5Expected {
		t.Fatalf("ProductService.Keyword is expected to equal keyword5_expected")
	}
}

func TestSetContentIDInProductService(t *testing.T) {
	srv := dummyProductService()

	var contentID = "15dss00145"
	srv.SetContentID(contentID)
	if srv.ContentID != contentID {
		t.Fatalf("ProductService.ContentID is expected to equal the input value. value:%s", contentID)
	}
}

func TestSetSiteInProductService(t *testing.T) {
	srv := dummyProductService()

	var site string

	site = SiteGeneral
	srv.SetSite(site)
	if srv.Site != site {
		t.Fatalf("ProductService.Site is expected to equal the input value. value:%s", site)
	}

	site = SiteAdult
	srv.SetSite(site)
	if srv.Site != site {
		t.Fatalf("ProductService.Site is expected to equal the input value. value:%s", site)
	}
}

func TestSetServiceInProductService(t *testing.T) {
	srv := dummyProductService()

	service := "digital"
	srv.SetService(service)
	if srv.Service != service {
		t.Fatalf("ProductService.Service is expected to equal the input value(service)")
	}
}

func TestSetFloorInProductService(t *testing.T) {
	srv := dummyProductService()

	floor := "videoa"
	srv.SetFloor(floor)
	if srv.Floor != floor {
		t.Fatalf("ProductService.Floor is expected to equal the input value(floor)")
	}
}

func TestSetGteDateInProductService(t *testing.T) {
	srv := dummyProductService()

	gte_date := "2016-04-01T00:00:00"
	srv.SetGteDate(gte_date)
	if srv.GteDate != gte_date {
		t.Fatalf("ProductService.GteDate is expected to equal the input value(gte_date)")
	}
}

func TestSetLteDateInProductService(t *testing.T) {
	srv := dummyProductService()

	lte_date := "2016-04-30T23:59:59"
	srv.SetGteDate(lte_date)
	if srv.GteDate != lte_date {
		t.Fatalf("ProductService.LteDate is expected to equal the input value(lte_date)")
	}
}

func TestValidateLengthInProductService(t *testing.T) {
	srv := dummyProductService()

	var target int64

	target = 1
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ProductService.ValidateLength is expected TRUE.")
	}

	target = DefaultProductAPILength
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ProductService.ValidateLength is expected TRUE.")
	}

	target = DefaultProductMaxLength
	srv.SetLength(target)
	if srv.ValidateLength() == false {
		t.Fatalf("ProductService.ValidateLength is expected TRUE.")
	}

	target = DefaultProductMaxLength + 1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ProductService.ValidateLength is expected FALSE.")
	}

	target = 0
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ProductService.ValidateLength is expected FALSE.")
	}

	target = -1
	srv.SetLength(target)
	if srv.ValidateLength() == true {
		t.Fatalf("ProductService.ValidateLength is expected FALSE.")
	}
}

func TestValidateOffsetInProductService(t *testing.T) {
	srv := dummyProductService()

	var target int64

	target = 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ProductService.ValidateOffset is expected TRUE. target:%d", target)
	}

	target = DefaultProductMaxOffset
	srv.SetOffset(target)
	if srv.ValidateOffset() == false {
		t.Fatalf("ProductService.ValidateOffset is expected TRUE. target:%d", target)
	}

	target = DefaultProductMaxOffset + 1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
	}

	target = 0
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
	}

	target = -1
	srv.SetOffset(target)
	if srv.ValidateOffset() == true {
		t.Fatalf("ProductService.ValidateOffset is expected FALSE. target:%d", target)
	}
}

func TestBuildRequestURLInProductService(t *testing.T) {
	var srv *ProductService
	var u string
	var err error
	var expected string

	srv = dummyProductService()
	srv.SetSite(SiteAdult)
	u, err = srv.BuildRequestURL()
	expected = APIBaseURL + "/ItemList?affiliate_id=" + DummyAffliateID + "&api_id=" + DummyAPIID + "&hits=" + strconv.FormatInt(DefaultProductAPILength, 10) + "&offset=" + strconv.FormatInt(DefaultAPIOffset, 10) + "&site=" + SiteAdult
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}

	srv = dummyProductService()
	srv.SetSite(SiteAdult)
	srv.SetLength(0)
	srv.SetOffset(0)
	u, err = srv.BuildRequestURL()
	expected = APIBaseURL + "/ItemList?affiliate_id=" + DummyAffliateID + "&api_id=" + DummyAPIID
	expectedBase := expected
	expected = expectedBase + "&site=" + SiteAdult
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}

	srv.SetSite("")
	u, err = srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ProductService.BuildRequestURL is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ProductService.BuildRequestURL is expected to return error.")
	}
	srv.SetSite(SiteAdult)

	srv.SetLength(-1)
	u, err = srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ProductService.BuildRequestURL is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ProductService.BuildRequestURL is expected to return error.")
	}
	srv.SetLength(0)

	srv.SetOffset(-1)
	u, err = srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ProductService.BuildRequestURL is expected empty if error occurs.")
	}
	if err == nil {
		t.Fatalf("ProductService.BuildRequestURL is expected to return error.")
	}
	srv.SetOffset(0)

	srv.SetSort("rank")
	expected = expectedBase + "&site=" + SiteAdult + "&sort=rank"
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ActressService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ActressService.BuildRequestURL is not expected to have error")
	}
	srv.SetSort("")

	srv.SetKeyword("上原亜衣")
	expected = expectedBase + "&keyword=" + url.QueryEscape("上原亜衣") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetKeyword("")

	srv.SetContentID("15dss00145")
	expected = expectedBase + "&cid=15dss00145&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetContentID("")

	srv.SetService("digital")
	expected = expectedBase + "&service=" + url.QueryEscape("digital") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetService("")

	srv.SetFloor("videoa")
	expected = expectedBase + "&floor=" + url.QueryEscape("videoa") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetFloor("")

	srv.SetArticle("actress")
	expected = expectedBase + "&article=" + url.QueryEscape("actress") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetArticle("")

	srv.SetArticleID("1011199")
	expected = expectedBase + "&article_id=" + url.QueryEscape("1011199") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetArticleID("")

	srv.SetStock("mono")
	expected = expectedBase + "&mono_stock=" + url.QueryEscape("mono") + "&site=" + SiteAdult
	u, err = srv.BuildRequestURL()
	if u != expected {
		t.Fatalf("ProductService.BuildRequestURL is expected to equal the expected value.\nexpected:%s\nactual:  %s", expected, u)
	}
	if err != nil {
		t.Fatalf("ProductService.BuildRequestURL is not expected to have error")
	}
	srv.SetStock("")
}

func TestBuildRequestURLWithoutApiIDInInProductService(t *testing.T) {
	srv := dummyProductService()
	srv.ApiID = ""
	u, err := srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ProductService.BuildRequestURL is expected empty if API ID is not set.")
	}
	if err == nil {
		t.Fatalf("ProductService.BuildRequestURL is expected to return error.")
	}
}

func TestBuildRequestURLWithWrongAffiliateIDInProductService(t *testing.T) {
	srv := dummyProductService()
	srv.AffiliateID = "fizzbizz-100"
	u, err := srv.BuildRequestURL()
	if u != "" {
		t.Fatalf("ProductService.BuildRequestURL is expected empty if wrong Affiliate ID is set.")
	}
	if err == nil {
		t.Fatalf("ProductService.BuildRequestURL is expected to return error.")
	}
}

func TestExcuteRequestProductAPIToServer(t *testing.T) {
	if !RequestAvailable {
		t.Skip("Not set valid credentials")
	}

	srv := NewProductService(TestAffiliateID, TestAPIID)
	srv.SetSite(SiteAdult)
	srv.SetService("mono")
	srv.SetFloor("dvd")
	srv.SetContentID("15dss00145")
	srv.SetSort("date")
	srv.SetLength(1)

	rst, err := srv.Execute()
	if err != nil {
		t.Skip("Maybe, The network is down.")
	}

	if reflect.TypeOf(rst).String() != "*api.ProductResponse" {
		t.Fatalf("ProductService.Execute is expected to return *api.ProductResponse")
	}
}

func dummyProductService() *ProductService {
	return NewProductService(DummyAffliateID, DummyAPIID)
}
