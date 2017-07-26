package api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type AuthorService struct {
	ApiID       string `mapstructure:"api_id"`
	AffiliateID string `mapstructure:"affiliate_id"`
	FloorID     string `mapstructure:"floor_id"`
	Initial     string `mapstructure:"initial"`
	Length      int64  `mapstructure:"hits"`
	Offset      int64  `mapstructure:"offset"`
}

type AuthorRawResponse struct {
	Request AuthorService  `mapstructure:"request"`
	Result  AuthorResponse `mapstructure:"result"`
}

type AuthorResponse struct {
	ResultCount   int64    `mapstructure:"result_count"`
	TotalCount    int64    `mapstructure:"total_count"`
	FirstPosition int64    `mapstructure:"first_position"`
	SiteName      string   `mapstructure:"site_name"`
	SiteCode      string   `mapstructure:"site_code"`
	ServiceName   string   `mapstructure:"service_name"`
	ServiceCode   string   `mapstructure:"service_code"`
	FloorID       string   `mapstructure:"floor_id"`
	FloorName     string   `mapstructure:"floor_name"`
	FloorCode     string   `mapstructure:"floor_code"`
	AuthorList    []Author `mapstructure:"author"`
}

type Author struct {
	AuthorID string `mapstructure:"author_id"`
	Name     string `mapstructure:"name"`
	Ruby     string `mapstructure:"ruby"`
	ListURL  string `mapstructure:"list_url"`
}

// NewAuthorService returns a new service for the given affiliate ID and API ID.
//
// NewAuthorServiceは渡したアフィリエイトIDとAPI IDを使用して新しい serviceを返します。
func NewAuthorService(affiliateID, apiID string) *AuthorService {
	return &AuthorService{
		ApiID:       apiID,
		AffiliateID: affiliateID,
		FloorID:     "",
		Initial:     "",
		Length:      DefaultAPILength,
		Offset:      DefaultAPIOffset,
	}
}

// Execute requests a url is created by BuildRequestURL.
// Use ExecuteWeak If you want get this response in interface{}.
//
// BuildRequestURLで生成したURLにリクエストします。
// もし interface{} でこのレスポンスを取得したい場合は ExecuteWeak を使用してください。
func (srv *AuthorService) Execute() (*AuthorResponse, error) {
	result, err := srv.ExecuteWeak()
	if err != nil {
		return nil, err
	}
	var raw AuthorRawResponse
	if err = mapstructure.WeakDecode(result, &raw); err != nil {
		return nil, err
	}
	return &raw.Result, nil
}

// ExecuteWeak requests a url is created by BuildRequestURL.
//
// BuildRequestURLで生成したURLにリクエストします。
func (srv *AuthorService) ExecuteWeak() (interface{}, error) {
	reqURL, err := srv.BuildRequestURL()
	if err != nil {
		return nil, err
	}

	return RequestJSON(reqURL)
}

// SetLength set the specified argument to AuthorService.Length
//
// SetLengthはLengthパラメータを設定します。
func (srv *AuthorService) SetLength(length int64) *AuthorService {
	srv.Length = length
	return srv
}

// SetHits sets the specified argument to AuthorService.Length
//  SetHits is the alias for SetLength
//
// SetHitsはLengthパラメータを設定します。
func (srv *AuthorService) SetHits(length int64) *AuthorService {
	srv.SetLength(length)
	return srv
}

// SetOffset sets the specified argument to AuthorService.Offset
//
// SetOffsetはOffsetパラメータを設定します。
func (srv *AuthorService) SetOffset(offset int64) *AuthorService {
	srv.Offset = offset
	return srv
}

// SetInitial sets the specified argument to AuthorService.Initial.
// This argment is author name's initial and you can use only hiragana.
//  e.g. srv.SetInitial("な") -> Soseki Natsume(なつめ そうせき, 夏目漱石)
//
// SetInitialはInitalパラメータに検索したい作者の頭文字をひらがなで設定します。
func (srv *AuthorService) SetInitial(initial string) *AuthorService {
	srv.Initial = TrimString(initial)
	return srv
}

// SetFloorID sets the specified argument to AuthorService.FloorID.
// You can retrieve Floor IDs from floor API.
//
// SetFloorIDはFloorIDパラメータを設定します。
// フロアIDはフロアAPIから取得できます。
func (srv *AuthorService) SetFloorID(floorID string) *AuthorService {
	srv.FloorID = TrimString(floorID)
	return srv
}

// ValidateLength validates AuthorService.Length within the range (1 <= value <= DefaultMaxLength).
// Refer to ValidateRange for more information about the range to validate.
//
// ValidateLengthはAuthorService.Lengthが範囲内(1 <= value <= DefaultMaxLength)にあるか検証します。
// 検証範囲について更に詳しく知りたい方はValidateRangeを参照してください。
func (srv *AuthorService) ValidateLength() bool {
	return ValidateRange(srv.Length, 1, DefaultMaxLength)
}

// ValidateOffset validates AuthorService.Offset within the range (1 <= value).
//
// ValidateOffsetはAuthorService.Offsetが範囲内(1 <= value)にあるか検証します。
func (srv *AuthorService) ValidateOffset() bool {
	return srv.Offset >= 1
}

// BuildRequestURL creates url to request author API.
//
// BuildRequestURLは作者検索APIにリクエストするためのURLを作成します。
func (srv *AuthorService) BuildRequestURL() (string, error) {
	if srv.ApiID == "" {
		return "", fmt.Errorf("set invalid ApiID parameter")
	}
	if !ValidateAffiliateID(srv.AffiliateID) {
		return "", fmt.Errorf("set invalid AffiliateID parameter")
	}
	if srv.FloorID == "" {
		return "", fmt.Errorf("set invalid FloorID parameter")
	}

	queries := url.Values{}
	queries.Set("api_id", srv.ApiID)
	queries.Set("affiliate_id", srv.AffiliateID)
	queries.Set("floor_id", srv.FloorID)

	if srv.Length != 0 {
		if !srv.ValidateLength() {
			return "", fmt.Errorf("length out of range: %d", srv.Length)
		}
		queries.Set("hits", strconv.FormatInt(srv.Length, 10))
	}

	if srv.Offset != 0 {
		if !srv.ValidateOffset() {
			return "", fmt.Errorf("offset out of range: %d", srv.Offset)
		}
		queries.Set("offset", strconv.FormatInt(srv.Offset, 10))
	}

	if srv.Initial != "" {
		queries.Set("initial", srv.Initial)
	}
	return APIBaseURL + "/AuthorSearch?" + queries.Encode(), nil
}
