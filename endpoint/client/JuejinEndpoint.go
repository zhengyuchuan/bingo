package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	httpendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"
	"net/url"
)

type JuejinRequest struct {
	Cursor   string `json:"cursor"`
	IdType   int    `json:"id_type"`
	Limit    int    `json:"limit"`
	SortType int    `json:"sort_type"`
}

type JuejinResponse struct {
	ErrNo   int    `json:"err_no"`
	ErrMsg  string `json:"err_msg"`
	Data    []Data `json:"data"`
	Cursor  string `json:"cursor"`
	Count   int    `json:"count"`
	HasMore bool   `json:"has_more"`
}
type MsgInfo struct {
	ID                int           `json:"id"`
	MsgID             string        `json:"msg_id"`
	UserID            string        `json:"user_id"`
	TopicID           string        `json:"topic_id"`
	Content           string        `json:"content"`
	PicList           []interface{} `json:"pic_list"`
	URL               string        `json:"url"`
	URLTitle          string        `json:"url_title"`
	URLPic            string        `json:"url_pic"`
	VerifyStatus      int           `json:"verify_status"`
	Status            int           `json:"status"`
	Ctime             string        `json:"ctime"`
	Mtime             string        `json:"mtime"`
	Rtime             string        `json:"rtime"`
	CommentCount      int           `json:"comment_count"`
	DiggCount         int           `json:"digg_count"`
	HotIndex          float64       `json:"hot_index"`
	RankIndex         float64       `json:"rank_index"`
	CommentScore      int           `json:"comment_score"`
	IsAdvertRecommend bool          `json:"is_advert_recommend"`
	AuditStatus       int           `json:"audit_status"`
}
type University struct {
	UniversityID string `json:"university_id"`
	Name         string `json:"name"`
	Logo         string `json:"logo"`
}
type Major struct {
	MajorID  string `json:"major_id"`
	ParentID string `json:"parent_id"`
	Name     string `json:"name"`
}
type ExtraMap struct {
}
type AuthorUserInfo struct {
	UserID                  string     `json:"user_id"`
	UserName                string     `json:"user_name"`
	Company                 string     `json:"company"`
	JobTitle                string     `json:"job_title"`
	AvatarLarge             string     `json:"avatar_large"`
	Level                   int        `json:"level"`
	Description             string     `json:"description"`
	FolloweeCount           int        `json:"followee_count"`
	FollowerCount           int        `json:"follower_count"`
	PostArticleCount        int        `json:"post_article_count"`
	DiggArticleCount        int        `json:"digg_article_count"`
	GotDiggCount            int        `json:"got_digg_count"`
	GotViewCount            int        `json:"got_view_count"`
	PostShortmsgCount       int        `json:"post_shortmsg_count"`
	DiggShortmsgCount       int        `json:"digg_shortmsg_count"`
	Isfollowed              bool       `json:"isfollowed"`
	FavorableAuthor         int        `json:"favorable_author"`
	Power                   int        `json:"power"`
	StudyPoint              int        `json:"study_point"`
	University              University `json:"university"`
	Major                   Major      `json:"major"`
	StudentStatus           int        `json:"student_status"`
	SelectEventCount        int        `json:"select_event_count"`
	SelectOnlineCourseCount int        `json:"select_online_course_count"`
	Identity                int        `json:"identity"`
	IsSelectAnnual          bool       `json:"is_select_annual"`
	SelectAnnualRank        int        `json:"select_annual_rank"`
	AnnualListType          int        `json:"annual_list_type"`
	ExtraMap                ExtraMap   `json:"extraMap"`
	IsLogout                int        `json:"is_logout"`
}
type Topic struct {
	TopicID       string `json:"topic_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Icon          string `json:"icon"`
	MsgCount      int    `json:"msg_count"`
	FollowerCount int    `json:"follower_count"`
	AttenderCount int    `json:"attender_count"`
}
type UserInteract struct {
	ID        int64 `json:"id"`
	Omitempty int   `json:"omitempty"`
	UserID    int64 `json:"user_id"`
	IsDigg    bool  `json:"is_digg"`
	IsFollow  bool  `json:"is_follow"`
	IsCollect bool  `json:"is_collect"`
}
type Org struct {
	OrgInfo    interface{} `json:"org_info"`
	OrgUser    interface{} `json:"org_user"`
	IsFollowed bool        `json:"is_followed"`
}
type Data struct {
	MsgID          string         `json:"msg_id"`
	MsgInfo        MsgInfo        `json:"msg_Info"`
	AuthorUserInfo AuthorUserInfo `json:"author_user_info"`
	Topic          Topic          `json:"topic"`
	UserInteract   UserInteract   `json:"user_interact"`
	Org            Org            `json:"org"`
}

func EncodeJuejin(cxt context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Cookie", "_ga=GA1.2.2010206452.1620609324; n_mh=g5uT9tzMjksycE0DN2xLd5pvdFboWpgOLRQtLQiy8Ds; passport_csrf_token_default=3dd15882c7ce7c6882c053534c2b3cdc; passport_csrf_token=3dd15882c7ce7c6882c053534c2b3cdc; sid_guard=609b8fb1965fb354caa2977cc6305229%7C1625794051%7C5184000%7CTue%2C+07-Sep-2021+01%3A27%3A31+GMT; uid_tt=d8517cfa25186846d36ec14145f9ce44; uid_tt_ss=d8517cfa25186846d36ec14145f9ce44; sid_tt=609b8fb1965fb354caa2977cc6305229; sessionid=609b8fb1965fb354caa2977cc6305229; sessionid_ss=609b8fb1965fb354caa2977cc6305229; MONITOR_WEB_ID=8634b3b5-ead5-47cc-b54c-fde4ed8d1f90; _gid=GA1.2.1348569692.1628472302")
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func DecodeJuejin(ctx context.Context, r *http.Response) (response interface{}, err error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp JuejinResponse
	err = json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func MakeJuejinEndpoint() httpendpoint.Endpoint {
	return httptransport.NewClient(
		"POST",
		getJuejinUrl(),
		EncodeJuejin,
		DecodeJuejin,
	).Endpoint()
}

func getJuejinUrl() *url.URL {
	juejinUrl, err := url.Parse("https://api.juejin.cn/recommend_api/v1/short_msg/recommend")
	if err != nil {
		return nil
	}
	return juejinUrl
}
