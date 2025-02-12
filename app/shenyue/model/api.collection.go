package model

// CollectionArticleReq 收藏文章请求结构体
type CollectionArticleReq struct {
	ArticleId int64 `json:"article_id"`
	Status    int64 `json:"status"` // 1 收藏 2 取消收藏
}

// CollectionArticleResp 收藏文章响应结构体
type CollectionArticleResp struct {
}

// CollectionArticleListReq 收藏文章列表请求结构体
type CollectionArticleListReq struct {
	UserId int64 `json:"user_id"`
}

// CollectionArticleListResp 收藏文章列表响应结构体
type CollectionArticleListResp struct {
	ArticleList []Article `json:"article_list"`
}
