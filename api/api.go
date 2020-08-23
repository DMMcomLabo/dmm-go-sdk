package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"
)

const (
	APIBaseURL       = "https://api.dmm.com/affiliate/v3"
	APIVersion       = "3"
	SiteGeneral      = "DMM.com"
	SiteAdult        = "FANZA"
	DefaultAPIOffset = 1
	DefaultAPILength = 100
	DefaultMaxLength = 500
)

// RequestJSON requests a retirived url and returns the response is parsed JSON-encoded data
//
// RequestJSONは指定されたURLにリクエストしJSONで返ってきたレスポンスをパースしたデータを返します。
func RequestJSON(url string) (interface{}, error) {
	// Ignore SSL Certificate Errors
	hc := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := hc.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error at API request:%#v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// TrimString wraps up strings.TrimString
//
// TrimStringはstrings.TrimStringのラップ関数です。
func TrimString(str string) string {
	return strings.TrimSpace(str)
}

// ValidateAffiliateID validates affiliate ID.
// (affiliate number range: 990 ~ 999)
//  e.g. dummy-999
//
// ValidateAffiliateIDはアフィリエイトID(例: dummy-999)のバリデーションを行います。
//（アフィリエイトの数値の範囲は 990〜999です）
func ValidateAffiliateID(affiliateID string) bool {
	if affiliateID == "" {
		return false
	}
	return regexp.MustCompile(`^.+-99[0-9]$`).Match([]byte(affiliateID))
}

// ValidateSite validates site parameter.
//
// ValidateSiteはsiteパラメータのバリデーションを行います
func ValidateSite(site string) bool {
	if site == "" {
		return false
	}
	if site != SiteGeneral && site != SiteAdult {
		return false
	}
	return true
}

// ValidateRange validates a retrieved number within the range ( number >= min && number <= max).
//
// ValidateRangeは指定された数値が最小値と最大値の範囲内にあるかどうか判定します。
func ValidateRange(target, min, max int64) bool {
	return target >= min && target <= max
}

// GetAPIVersion returns API version.
//
// GetAPIVersionはAPIのバージョンを返します。
func GetAPIVersion() string {
	return APIVersion
}

// buildAPIEndpoint returns API Endpoint path.
//
// buildAPIEndpointはAPIエンドポイントのフルパスを組み立てて返します。
func buildAPIEndpoint(p string) (*url.URL, error) {
	u, err := url.Parse(APIBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Parse error: %#v", err)
	}
	u.Path = path.Join(u.Path, p)
	return u, err
}
