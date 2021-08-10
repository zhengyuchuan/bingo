package service

import (
	"context"
	httpendpoint "github.com/go-kit/kit/endpoint"
)

type SvcRequest struct {
	Name string `http:"name" json:"name"`
	Age  string `http:"age" json:"age"`
}

type SvcResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Data []Data `json:"data"`
}

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
	HotIndex          int           `json:"hot_index"`
	RankIndex         int           `json:"rank_index"`
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

func SvcService(ctx context.Context, svc SvcRequest, ed httpendpoint.Endpoint) SvcResponse {
	data, err := ed(ctx, JuejinRequest{
		Cursor:   "0",
		IdType:   4,
		Limit:    20,
		SortType: 300,
	})
	if err != nil {
		return SvcResponse{
			Id:   0,
			Name: "zheng",
			Age:  10,
			Data: nil,
		}
	}
	newData, ok := data.(JuejinResponse)
	if !ok {
		return SvcResponse{
			Id:   0,
			Name: "zheng",
			Age:  10,
			Data: nil,
		}
	}

	return SvcResponse{
		Id:   0,
		Name: "zheng",
		Age:  10,
		Data: newData.Data,
	}
}
