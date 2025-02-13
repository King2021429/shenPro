package model

// LikeArticleReq 点赞请求结构体
type LikeArticleReq struct {
	ArticleId int64 `json:"article_id"`
	Status    int64 `json:"status"` // 1 点赞 2 点踩
}

// LikeArticleResp 点赞响应结构体
type LikeArticleResp struct {
}

// LikeArticleListReq 点赞列表请求结构体
type LikeArticleListReq struct {
	Uid int64 `json:"uid"`
}

type LikeArticleListResp struct {
	ArticleList []Article `json:"article_list"`
}
