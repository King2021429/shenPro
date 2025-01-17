package model

// LikeArticleReq 点赞请求结构体
type LikeArticleReq struct {
	ArticleId int64 `json:"article_id"`
	Status    int64 `json:"status"` // 1 点赞 2 点踩
}

// LikeArticleResp 点赞响应结构体
type LikeArticleResp struct {
}
